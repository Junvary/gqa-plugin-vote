<template>
    <q-page padding>
        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.voteMonth" label="选择年月" clearable placeholder="202201" />
            <q-select style="width: 20%" v-model="queryParams.voteType" :options="dictOptions.voteType" emit-value
                map-options label="投票类型" @update:model-value="getTableDataAll" />
            <q-btn color="primary" @click="getTableDataAll" label="确定" />
        </div>
        <q-card style="margin: 15px 0" v-if="tableDataVoted.length">
            <div class="row justify-center text-h6">
                本期已有{{ tableDataVoted.length }}人完成投票
            </div>
            <div class="q-guuter-md">
                <q-chip dense class="glossy" color="positive" text-color="white" v-for="item in tableDataVoted"
                    :key="item.voteFrom">
                    <GqaShowName :customNameObject="item.voteFromByUser" />
                    ({{ item.voteFrom }})
                </q-chip>
            </div>
        </q-card>

        <q-separator />

        <div ref="monthscorechart" style="width: 100%; height: 500px"></div>
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <span class="text-h6">
                    <!-- {{ queryParams.voteMonth }}
                    <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" /> -->
                    {{ chartTitle }}
                    投票明细
                </span>
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-voteType="props">
                <q-td :props="props">
                    <GqaDictShow dictName="voteType" :dictCode="props.row.voteType" />
                </q-td>
            </template>

            <template v-slot:body-cell-voteTypeDetail="props">
                <q-td :props="props" v-if="queryParams.voteType === 'dy'">
                    <GqaDictShow dictName="voteTypeDetailDy" :dictCode="props.row.voteTypeDetail" />
                </q-td>
                <q-td :props="props" v-if="queryParams.voteType === 'gl'">
                    <GqaDictShow dictName="voteTypeDetailGl" :dictCode="props.row.voteTypeDetail" />
                </q-td>
            </template>

            <template v-slot:body-cell-candidate="props">
                <q-td :props="props">
                    <GqaShowName :customNameObject="props.row.candidateByUser" />
                    ({{ props.row.candidate }})
                </q-td>
            </template>

            <template v-slot:body-cell-voteRatio="props">
                <q-td :props="props" v-if="queryParams.voteType === 'dy'">
                    <GqaDictShow dictName="voteDyRatio" :dictCode="props.row.voteRatio" />
                    {{ props.row.ratio }}%
                </q-td>
                <q-td :props="props" v-if="queryParams.voteType === 'gl'">
                    <GqaDictShow dictName="voteGlRatio" :dictCode="props.row.voteRatio" />
                    {{ props.row.ratio }}%
                </q-td>
            </template>

            <template v-slot:body-cell-voteFrom="props">
                <q-td :props="props">
                    <GqaShowName :customNameObject="props.row.voteFromByUser" />
                    ({{ props.row.voteFrom }})
                </q-td>
            </template>

            <template v-slot:body-cell-createdAt="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.createdAt) }}
                </q-td>
            </template>

        </q-table>
    </q-page>
</template>


<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar, date } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref, markRaw } from 'vue'
import { useI18n } from 'vue-i18n'
const echarts = require('echarts')

const monthscorechart = ref(null)
const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'plugin-vote/vote-result-list',
    chart: 'plugin-vote/vote-result-chart',
    voted: 'plugin-vote/vote-result-voted',
}
const columns = computed(() => {
    return [
        { name: 'voteType', align: 'center', label: '投票类型', field: 'voteType' },
        { name: 'voteTypeDetail', align: 'center', label: '投票细类', field: 'voteTypeDetail' },
        { name: 'candidate', align: 'center', label: '候选人', field: 'candidate' },
        { name: 'voteScore', align: 'center', label: '投票分数', field: 'voteScore' },
        { name: 'voteRatio', align: 'center', label: '投票权重', field: 'voteRatio' },
        { name: 'voteFrom', align: 'center', label: '投票人', field: 'voteFrom' },
        { name: 'voteMonth', align: 'center', label: '投票周期', field: 'voteMonth' },
        { name: 'createdAt', align: 'center', label: '投票时间', field: 'createdAt' },
    ]
})
const {
    dictOptions,
    showDateTime,
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
    GqaShowName,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    onRequest,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

const tableDataVoted = ref([])
const tableDataChart = ref([])
const myChart = ref(null)
onMounted(() => {
    let ct = monthscorechart.value;
    echarts.dispose(ct);
    myChart.value = markRaw(echarts.init(ct));

    const timeStamp = Date.now()
    queryParams.value.voteMonth = date.formatDate(timeStamp, 'YYYYMM')
    queryParams.value.voteType = 'dy'
    pagination.value.sortBy = 'created_at'
    getTableDataAll()
})

const options = ref({})
const monthScore = ref({
    user: [],
})
const monthUserScore = ref([])
const paginationChart = ref({
    sortBy: 'created_at',
    descending: true,
    page: 1,
    rowsPerPage: 999999,
})
const chartTitle = ref('')
const getTableDataAll = () => {
    getTableDataList()
    getTableDataChart()
    getTableDataVoted()
}
const getTableDataList = () => {
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    })
}
const getTableDataVoted = async () => {
    await postAction(url.voted, queryParams.value).then((res) => {
        if (res.code === 1) {
            tableDataVoted.value = res.data.records
        }
    })
}
const getTableDataChart = async () => {
    await onRequestChart({ pagination: paginationChart.value })
}
const onRequestChart = async (props) => {
    paginationChart.value = []
    // 组装分页和过滤条件
    const params = {}
    params.sortBy = props.pagination.sortBy
    params.desc = props.pagination.descending
    params.page = props.pagination.page
    params.pageSize = props.pagination.rowsPerPage
    const allParams = Object.assign({}, params, queryParams.value)
    // 带参数请求数据
    await postAction(url.chart, allParams).then((res) => {
        if (res.code === 1) {
            // 最终要把分页给同步掉
            paginationChart.value = props.pagination
            // 并且加入总数字段
            paginationChart.value.rowsNumber = res.data.total
            tableDataChart.value = res.data.records
            changeTableData()
        }
    })
}
const changeTableData = () => {
    changeChartTitle()
    for (let i of tableDataChart.value) {
        i.candidateMonth = i.candidate + i.voteMonth
    }
    monthScore.value = {
        user: [],
    }
    monthUserScore.value = []
    getMonthScore(queryParams.value.voteMonth)
}
const changeChartTitle = () => {
    if (queryParams.value.voteMonth) {
        chartTitle.value = dictOptions.value.voteType.filter((item) => item.dict_code === queryParams.value.voteType)[0].dict_label + queryParams.value.voteMonth
    } else {
        chartTitle.value = dictOptions.value.voteType.filter((item) => item.dict_code === queryParams.value.voteType)[0].dict_label + '(总)'
    }
}
const getMonthScore = (month) => {
    const monthScoreTemp = tableDataChart.value
    const onlyIdList = []
    for (let i of monthScoreTemp) {
        if (onlyIdList.indexOf(i.candidateMonth) === -1) {
            onlyIdList.push(i.candidateMonth)
        }
    }
    for (let id of onlyIdList) {
        const oneCandidate = monthScoreTemp.filter((item) => item.candidateMonth === id)
        if (oneCandidate.length) {
            monthScore.value.user.push(oneCandidate[0].candidateByUser.real_name)
            const detailList = []
            for (let v of oneCandidate) {
                monthScore.value[v.voteTypeDetail] = monthScore.value[v.voteTypeDetail] ? monthScore.value[v.voteTypeDetail] : []
                if (detailList.indexOf(v.voteTypeDetail) === -1) {
                    detailList.push(v.voteTypeDetail)
                }
            }
            for (let d of detailList) {
                const thisDetailList = oneCandidate.filter((item) => item.voteTypeDetail === d)
                let total = 0
                for (let t of thisDetailList) {
                    total += t.voteScore
                }
                monthScore.value[d].push((total / thisDetailList.length).toFixed(2))
            }
        }
    }

    // 未选择quearyParams.voteMonth的情况，展示所有分数，需要将重复人合并处理
    if (!queryParams.value.voteMonth) {
        const allScore = {
            user: [],
        }
        for (let i in monthScore.value.user) {
            if (allScore.user.indexOf(monthScore.value.user[i]) === -1) {
                allScore.user.push(monthScore.value.user[i])
                for (let item in monthScore.value) {
                    if (item !== 'user') {
                        if (allScore[item] && allScore[item].length) {
                            allScore[item].push(monthScore.value[item][i])
                        } else {
                            allScore[item] = [monthScore.value[item][i]]
                        }
                    }
                }
            } else {
                for (let item in monthScore.value) {
                    const allScoreDetailKey = allScore.user.indexOf(monthScore.value.user[i])
                    if (item !== 'user') {
                        allScore[item][allScoreDetailKey] = Number(allScore[item][allScoreDetailKey]) + Number(monthScore.value[item][i])
                    }
                }
            }
        }
        for (let userIndex in allScore.user) {
            for (let i in allScore) {
                if (i !== 'user') {
                    allScore[i][userIndex] = allScore[i][userIndex] / monthScore.value.user.filter((item) => item === allScore.user[userIndex]).length
                }
            }
        }

        // 计算平均分
        monthScore.value = allScore
    }

    // 计算平均分，管理人员
    if (queryParams.value.voteType === 'gl') {
        for (let u = 0; u < monthScore.value.user.length; u++) {
            let userScore = 0
            let userNumber = 1
            for (let i in monthScore.value) {
                if (i !== 'user') {
                    userNumber += 1
                    userScore += Number(monthScore.value[i][u])
                }
            }
            monthUserScore.value.push((userScore / (userNumber - 1)).toFixed(2))
        }
    }

    // 计算平均分，党员投票需要算出前五项的平均值，再加上两项额外的评价分，最终得总分
    if (queryParams.value.voteType === 'dy') {
        for (let u = 0; u < monthScore.value.user.length; u++) {
            let userScore = 0
            let userNumber = 1
            let pingjiaScore = 0
            for (let i in monthScore.value) {
                if (i !== 'user' && i !== 'dy_p_jijian' && i !== 'dy_p_zhenggong') {
                    userNumber += 1
                    userScore += Number(monthScore.value[i][u])
                }
                if (i === 'dy_p_jijian' || i === 'dy_p_zhenggong') {
                    pingjiaScore += Number(monthScore.value[i][u])
                }
            }
            monthUserScore.value.push((userScore / (userNumber - 1) + pingjiaScore).toFixed(2))
        }
    }
    updateMonthScoreEcharts()
}
const updateMonthScoreEcharts = () => {
    const series = []
    if (queryParams.value.voteType === 'dy') {
        for (let dict of dictOptions.value.voteTypeDetailDy) {
            series.push({
                name: dict.dictLabel,
                type: 'bar',
                data: monthScore.value[dict.dict_code],
            })
        }
    }
    if (queryParams.value.voteType === 'gl') {
        for (let dict of dictOptions.value.voteTypeDetailGl) {
            series.push({
                name: dict.dictLabel,
                type: 'bar',
                data: monthScore.value[dict.dict_code],
            })
        }
    }

    series.push({
        name: '总得分',
        type: 'line',
        yAxisIndex: 1,
        data: monthUserScore.value,
        label: {
            show: true,
            formatter: '{c}',
        },
    })
    options.value = {
        title: {
            text: chartTitle.value + ' 得分情况',
        },
        tooltip: {
            trigger: 'item',
        },
        grid: {
            left: '3%',
            right: '3%',
            bottom: '3%',
            containLabel: true,
        },
        toolbox: {
            feature: {
                dataView: { show: true, readOnly: false },
                magicType: { show: true, type: ['line', 'bar'] },
                restore: { show: true },
                saveAsImage: { show: true },
            },
        },
        legend: {},
        xAxis: [
            {
                type: 'category',
                data: monthScore.value.user,
                axisPointer: {
                    type: 'shadow',
                },
                axisLabel: {
                    interval: 0,
                    rotate: 40,
                },
            },
        ],
        yAxis: [
            {
                type: 'value',
                name: '分项平均得分',
                axisLabel: {
                    formatter: '{value}',
                },
            },
            {
                type: 'value',
                name: '总得分',
                axisLabel: {
                    formatter: '{value}',
                },
            },
        ],
        series: series,
    }
    myChart.value.clear()
    myChart.value.setOption(options.value)
}
</script>
