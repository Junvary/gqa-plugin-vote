<template>
    <q-page padding>
        <div class="row">
            <div class="col" style="padding: 4px">
                <q-toolbar class="bg-primary text-white shadow-1">
                    <q-toolbar-title>民主评议业务骨干及管理人员投票</q-toolbar-title>
                    <q-space></q-space>
                    <span v-if="canVoteGl && candidateListGl.length">
                        共:
                        {{ candidateListGl.length }}
                        人&nbsp;&nbsp;&nbsp;&nbsp;
                    </span>
                    <q-btn push glossy color="negative" @click="handleVoteGl"
                        v-if="canVoteGl && candidateListGl.length">
                        提交本次投票结果
                    </q-btn>
                    <span v-else-if="!candidateListGl.length">
                        暂时没有候选人
                    </span>
                    <span v-else>
                        本轮已投票或没有投票权限!
                    </span>
                </q-toolbar>
                <q-form ref="candidateFormGl" v-if="canVoteGl" class="gqa-form">
                    <template v-if="candidateListGl.length">
                        <q-card v-for="(item, index) in candidateListGl" :key="index" bordered>
                            <q-card-section style="padding: 0 4px">
                                <div class="row justify-between items-center">
                                    <div class="col-3 row justify-center items-center">
                                        <q-chip class="glossy" color="primary" text-color="white">
                                            <GqaShowName :customNameObject="item.candidateByUser" />
                                            ({{ item.candidate }})
                                            : {{ personTotalScore(voteResultGl[item.candidate]) }}
                                        </q-chip>
                                    </div>
                                    <div class="col-9 row">
                                        <q-input class="col" v-for="(dict, index) in dictOptions.voteTypeDetailGl"
                                            :key="index" v-model.number="voteResultGl[item.candidate][dict.dict_code]"
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
const personTotalScore = computed(() => {
    return (scoreList) => {
        let t = 0
        for (let i in scoreList) {
            t += scoreList[i]
        }
        return t
    }
})
const url = {
    list: 'plugin-vote/candidate-list',
    vote: 'plugin-vote/vote-handle',
    canVote: 'plugin-vote/can-vote-gl',
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

onMounted(() => {
    pagination.value.sortBy = 'candidate'
    pagination.value.rowsPerPage = 9999999
    checkCanVote()
    changeTableData()
})
const candidateListGl = ref([])
const voteResultGl = ref({})
const canVoteGl = ref(false)

const checkCanVote = () => {
    postAction(url.canVote).then((res) => {
        if (res.code === 1) {
            canVoteGl.value = res.data.records.gl
        }
    })
}
const changeTableData = () => {
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    }).then(() => {
        candidateListGl.value = tableData.value.filter((item) => item.voteType === 'gl')
        for (let gl of candidateListGl.value) {
            for (let dict of dictOptions.value.voteTypeDetailGl) {
                voteResultGl.value[gl.candidate] = Object.assign({}, voteResultGl.value[gl.candidate], {
                    [dict.dict_code]: 95,
                })
            }
        }
    })
}
const candidateFormGl = ref(null)
const handleVoteGl = async () => {
    const success = await candidateFormGl.value.validate()
    if (success) {
        $q.dialog({
            title: '民主评议业务骨干及管理人员投票',
            message: '确定提交本次民主评议业务骨干及管理人员投票结果吗？',
            cancel: true,
            persistent: true,
        }).onOk(async () => {
            const voteList = []
            for (let c in voteResultGl.value) {
                for (let r in voteResultGl.value[c]) {
                    voteList.push({
                        candidate: c,
                        voteType: 'gl',
                        voteTypeDetail: r,
                        voteScore: voteResultGl.value[c][r],
                    })
                }
            }
            handleVotePost(voteList, '民主评议业务骨干及管理人员投票')
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
