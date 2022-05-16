<template>
    <q-page padding>
        <div class="row">
            <div class="col" style="padding: 4px">
                <q-toolbar class="bg-primary text-white shadow-1">
                    <q-toolbar-title>
                        民主评议党员投票
                        <span v-if="canVoteJiJian">
                            & 纪检委员评价
                        </span>
                        <span v-if="canVoteZhengGong">
                            & 政工干事评价
                        </span>
                    </q-toolbar-title>
                    <q-space></q-space>
                    <span v-if="canVoteDy && candidateListDy.length">
                        共:
                        {{ candidateListDy.length }}
                        人&nbsp;&nbsp;&nbsp;&nbsp;
                    </span>
                    <q-btn push glossy color="negative" @click="handleVoteDy"
                        v-if="canVoteDy && candidateListDy.length">
                        提交本次投票结果
                    </q-btn>
                    <span v-else-if="!candidateListDy.length">
                        暂时没有候选人
                    </span>
                    <span v-else>
                        本轮已投票或没有投票权限!
                    </span>
                </q-toolbar>
                <q-form ref="candidateFormDy" v-if="canVoteDy" class="gqa-form">
                    <template v-if="candidateListDy.length">
                        <q-card v-for="(item, index) in candidateListDy" :key="index" bordered>
                            <q-card-section style="padding: 0 4px">
                                <div class="row justify-between items-center">
                                    <div class="col-3 row justify-center items-center">
                                        <q-chip class="glossy" color="primary" text-color="white">
                                            <GqaShowName :customNameObject="item.candidateByUser" />
                                            ({{ item.candidate }})
                                            : {{ personTotalScore(voteResultDy[item.candidate]) }}分
                                        </q-chip>
                                    </div>
                                    <div class="col-9 row">
                                        <q-input class="col" v-for="(dict, index) in trueVoteTypeDetailDy" :key="index"
                                            v-model.number="voteResultDy[item.candidate][dict.dict_code]"
                                            :label="dict.dict_label" type="number"
                                            :rules="[val => val >= 1 && val <= 100 || '必须大于等于1，且小于等于100']" />
                                    </div>
                                </div>
                            </q-card-section>
                        </q-card>
                    </template>
                    <span v-else>
                        暂时没有候选人!
                    </span>
                </q-form>
            </div>
        </div>
    </q-page>
</template>

<script setup>
import { useQuasar } from 'quasar';
import useTableData from 'src/composables/useTableData'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n';
import { postAction } from 'src/api/manage';

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'plugin-vote/candidate-list',
    vote: 'plugin-vote/vote-handle',
    canVote: 'plugin-vote/can-vote-dy',
}
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

const personTotalScore = computed(() => {
    return (scoreList) => {
        let t = 0
        for (let i in scoreList) {
            t += scoreList[i]
        }
        return t
    }
})
onMounted(() => {
    pagination.value.sortBy = 'candidate'
    pagination.value.rowsPerPage = 9999999
    checkCanVote()
})

const candidateListDy = ref([])
const voteResultDy = ref({})
const canVoteDy = ref(false)
const canVoteJiJian = ref(false)
const canVoteZhengGong = ref(false)
const trueVoteTypeDetailDy = ref([])

const checkCanVote = () => {
    postAction(url.canVote).then((res) => {
        if (res.code === 1) {
            canVoteDy.value = res.data.records.dy
            canVoteJiJian.value = res.data.records.jiJian
            canVoteZhengGong.value = res.data.records.zhengGong
            if (canVoteDy.value && !canVoteJiJian.value && !canVoteZhengGong.value) {
                trueVoteTypeDetailDy.value = dictOptions.value.voteTypeDetailDy.filter((item) => item.dict_code !== 'dy_p_zhenggong' && item.dict_code !== 'dy_p_jijian')
            } else if (canVoteDy.value && !canVoteJiJian.value && canVoteZhengGong.value) {
                trueVoteTypeDetailDy.value = dictOptions.value.voteTypeDetailDy.filter((item) => item.dict_code === 'dy_p_zhenggong')
            } else if (canVoteDy.value && canVoteJiJian.value && !canVoteZhengGong.value) {
                trueVoteTypeDetailDy.value = dictOptions.value.voteTypeDetailDy.filter((item) => item.dict_code === 'dy_p_jijian')
            } else if (canVoteDy.value && canVoteJiJian.value && canVoteZhengGong.value) {
                trueVoteTypeDetailDy.value = dictOptions.value.voteTypeDetailDy.filter((item) => item.dict_code === 'dy_p_zhenggong' || item.dict_code === 'dy_p_jijian')
            } else {
                trueVoteTypeDetailDy.value = []
            }
            changeTableData()
        }
    })
}
const changeTableData = () => {
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    }).then(() => {
        candidateListDy.value = tableData.value.filter((item) => item.voteType === 'dy')
        for (let dy of candidateListDy.value) {
            for (let dict of trueVoteTypeDetailDy.value) {
                voteResultDy.value[dy.candidate] = Object.assign({}, voteResultDy.value[dy.candidate], {
                    [dict.dict_code]: 95,
                })
            }
        }
    })
}
const candidateFormDy = ref(null)
const handleVoteDy = async () => {
    const success = await candidateFormDy.value.validate()
    if (success) {
        $q.dialog({
            title: '民主评议党员投票',
            message: '确定提交本次民主评议党员投票结果吗？',
            cancel: true,
            persistent: true,
        }).onOk(async () => {
            const voteList = []
            for (let c in voteResultDy.value) {
                for (let r in voteResultDy.value[c]) {
                    voteList.push({
                        candidate: c,
                        voteType: 'dy',
                        voteTypeDetail: r,
                        voteScore: voteResultDy.value[c][r],
                    })
                }
            }
            handleVotePost(voteList, '民主评议党员投票')
        })
    } else {
        $q.notify({
            type: 'negative',
            message: t('FixForm'),
        })
    }
}
const handleVotePost = (voteList, type) => {
    postAction(url.vote, {
        requestAddScoreDetail: voteList,
    }).then((res) => {
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: '投票成功: ' + type,
            })
            checkCanVote()
        }
    })
}
</script>
