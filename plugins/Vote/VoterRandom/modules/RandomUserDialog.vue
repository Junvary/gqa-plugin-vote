<template>
    <q-dialog v-model="randomUserVisible" position="top">
        <q-card style="width: 80vw; max-width: 80vw;">
            <div class="row">
                <q-card-section align="center" class="col">
                    <q-table row-key="id" separator="cell" :rows="tableDataBase" :columns="columns"
                        v-model:pagination="paginationBase" :rows-per-page-options="[0]" :loading="loadingBase"
                        class="my-sticky-header-table" @request="onRequestBase">
                        <template v-slot:top>
                            <span class="row text-h6">
                                <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                                : 固定投票人
                            </span>
                            <q-space></q-space>
                            <q-btn label="仅使用固定投票人" color="primary" @click="handleSaveRandom('base')" />
                        </template>
                        <template v-slot:body-cell-avatar="props">
                            <q-td :props="props">
                                <GqaAvatar :src="props.row.voterByUser.avatar" />
                            </q-td>
                        </template>

                        <template v-slot:body-cell-username="props">
                            <q-td :props="props">
                                {{ props.row.voter }}
                            </q-td>
                        </template>

                        <template v-slot:body-cell-nickname="props">
                            <q-td :props="props">
                                {{ props.row.voterByUser.nickname }}
                            </q-td>
                        </template>

                        <template v-slot:body-cell-realName="props">
                            <q-td :props="props">
                                {{ props.row.voterByUser.realName }}
                            </q-td>
                        </template>

                        <template v-slot:body-cell-voteType="props">
                            <q-td :props="props">
                                <GqaDictShow dictName="voteType" :dictCode="props.row.voteType" />
                            </q-td>
                        </template>
                        <template v-slot:body-cell-voteRatio="props">
                            <q-td :props="props" v-if="props.row.voteType === 'dy'">
                                <GqaDictShow dictName="voteDyRatio" :dictCode="props.row.voteRatio" withExt1 ext1="%" />
                            </q-td>
                            <q-td :props="props" v-if="props.row.voteType === 'gl'">
                                <GqaDictShow dictName="voteGlRatio" :dictCode="props.row.voteRatio" withExt1 ext1="%" />
                            </q-td>
                        </template>
                    </q-table>
                </q-card-section>
                <q-card-section align="center" class="col">
                    <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns"
                        v-model:pagination="pagination" :rows-per-page-options="[0]"
                        :class="tableData.length ? 'my-sticky-header-table2' : ''" :loading="loading"
                        @request="onRequest">
                        <template v-slot:top>
                            <span class="row text-h6">
                                <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                                : 随机投票人
                            </span>
                            <q-space></q-space>
                        </template>
                        <template v-slot:body-cell-avatar="props">
                            <q-td :props="props">
                                <GqaAvatar :src="props.row.avatar" />
                            </q-td>
                        </template>

                        <template v-slot:body-cell-username="props">
                            <q-td :props="props">
                                {{ props.row.username }}
                            </q-td>
                        </template>

                        <template v-slot:body-cell-nickname="props">
                            <q-td :props="props">
                                {{ props.row.nickname }}
                            </q-td>
                        </template>

                        <template v-slot:body-cell-realName="props">
                            <q-td :props="props">
                                {{ props.row.realName }}
                            </q-td>
                        </template>

                        <template v-slot:body-cell-voteType="props">
                            <q-td :props="props">
                                <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                            </q-td>
                        </template>

                        <template v-slot:body-cell-voteRatio="props">
                            <q-td :props="props" v-if="queryParams.voteType === 'dy'">
                                其他评议人员 40 %
                            </q-td>
                            <q-td :props="props" v-if="queryParams.voteType === 'gl'">
                                职工代表/普通职工 30%
                            </q-td>
                        </template>

                    </q-table>

                    <q-form ref="randomUserForm">
                        <q-input v-model.number="queryParams.randomNumber" type="number"
                            :rules="[val => val > 0 || '随机投票人数量必须大于0才能抽取']" label="选择随机投票人数量" />
                    </q-form>
                    <div class="q-gutter-xs">
                        <q-btn label="选取随机投票人" color="primary" @click="handleRandom" />
                        <q-btn v-if="tableData.length" label="确认选取固定+随机投票人" color="negative"
                            @click="handleSaveRandom('all')" />
                    </div>
                </q-card-section>
            </div>
            <q-form ref="randomMemoForm" style="margin: 0 20px">
                <q-input v-model="memo" label="投票说明/备注/版本" :rules="[val => val && val.length > 0 || $t('NeedInput')]"
                    placeholder="xxxx年第xx期xx投票" />
            </q-form>

            <!-- <q-separator /> -->

            <div class="row justify-center items-center q-gutter-xs" style="margin: 10px 0">
                <q-btn label="取消" color="negative" v-close-popup />
            </div>

        </q-card>
    </q-dialog>
</template>


<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    baseList: 'plugin-vote/voter-list',
    list: 'plugin-vote/voter-random-get',
    add: 'plugin-vote/voter-random-add',
}
const columns = computed(() => {
    return [
        { name: 'avatar', align: 'center', label: t('Avatar'), field: 'avatar' },
        { name: 'username', align: 'center', label: t('Username'), field: 'username' },
        { name: 'nickname', align: 'center', label: t('Nickname'), field: 'nickname' },
        { name: 'realName', align: 'center', label: t('RealName'), field: 'realName' },
        { name: 'voteType', align: 'center', label: '投票类型', field: 'voteType' },
        { name: 'voteRatio', align: 'center', label: '投票权重', field: 'voteRatio' },
    ]
})
const {
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
    getTableData,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    getTableData()
})
const loadingBase = ref(false)
const tableDataBase = ref([])
const paginationBase = ref({
    sortBy: 'sort',
    descending: false,
    page: 1,
    rowsPerPage: 20000,
})
const randomUserVisible = ref(false)
const memo = ref('')
const show = (voteType) => {
    pagination.value.rowsPerPage = 9999999
    queryParams.value.randomNumber = 0
    queryParams.value.voteType = voteType
    queryParams.value.randomNumber = 0
    tableData.value = []
    randomUserVisible.value = true
    getTableDataBase()
}
defineExpose({
    show
})
const getTableDataBase = async () => {
    await onRequestBase({ pagination: paginationBase.value })
}
const onRequestBase = async (props) => {
    tableDataBase.value = []
    loadingBase.value = true
    // 组装分页和过滤条件
    const params = {}
    params.sortBy = props.pagination.sortBy
    params.desc = props.pagination.descending
    params.page = props.pagination.page
    params.pageSize = props.pagination.rowsPerPage
    const allParams = Object.assign({}, params, queryParams.value)
    // 带参数请求数据
    await postAction(url.baseList, allParams).then((res) => {
        if (res.code === 1) {
            // 最终要把分页给同步掉
            paginationBase.value = props.pagination
            // 并且加入总数字段
            paginationBase.value.rowsNumber = res.data.total
            tableDataBase.value = res.data.records
        }
    }).finally(() => {
        loadingBase.value = false
    })
}
const randomUserForm = ref(null)
const handleRandom = async () => {
    const success = await randomUserForm.value.validate()
    if (success) {
        getTableData()
    } else {
        $q.notify({
            type: 'negative',
            message: t('FixForm'),
        })
    }
}
const randomMemoForm = ref(null)
const emit = defineEmits(['raomdomOver'])
const handleSaveRandom = async (type) => {
    const success = await randomMemoForm.value.validate()
    if (success) {
        const voterList = []
        for (let item of tableDataBase.value) {
            voterList.push(item.voter)
        }
        if (type === 'all') {
            for (let r of tableData.value) {
                voterList.push(r.username)
            }
        }
        postAction(url.add, {
            voteType: queryParams.value.voteType,
            memo: memo.value,
            voter: voterList,
        }).then((res) => {
            if (res.code === 1) {
                $q.notify({
                    type: 'positive',
                    message: '创建新版本投票人列表成功!',
                })
            }
            emit('raomdomOver')
            randomUserVisible.value = false
        })
    } else {
        $q.notify({
            type: 'negative',
            message: t('FixForm'),
        })
    }
}
</script>

<style lang="scss" scoped>
.my-sticky-header-table {
    height: 75vh;

    thead tr th {
        position: sticky;
        z-index: 1;
    }

    thead tr:first-child th {
        top: 10;
    }
}

.my-sticky-header-table2 {
    height: 65vh;

    thead tr th {
        position: sticky;
        z-index: 1;
    }

    thead tr:first-child th {
        top: 10;
    }
}
</style>