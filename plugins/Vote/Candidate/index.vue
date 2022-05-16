<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.candidate" label="候选人" clearable />
            <q-select style="width: 20%" v-model="queryParams.voteType" :options="dictOptions.voteType" emit-value
                map-options label="投票类型" @update:model-value="handleSearch" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn dense color="primary" @click="showAddUserForm()">
                    新增候选人(
                    <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                    )
                </q-btn>
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-avatar="props">
                <q-td :props="props">
                    <GqaAvatar :src="props.row.candidateByUser.avatar" />
                </q-td>
            </template>

            <template v-slot:body-cell-username="props">
                <q-td :props="props">
                    {{ props.row.candidate }}
                </q-td>
            </template>

            <template v-slot:body-cell-nickname="props">
                <q-td :props="props">
                    {{ props.row.candidateByUser.nickname }}
                </q-td>
            </template>

            <template v-slot:body-cell-realName="props">
                <q-td :props="props">
                    {{ props.row.candidateByUser.realName }}
                </q-td>
            </template>

            <template v-slot:body-cell-voteType="props">
                <q-td :props="props">
                    <GqaDictShow dictName="voteType" :dictCode="props.row.voteType" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>

        <SelectUserDialog ref="selectUserDialog" @handleSelectUser="handleAddUser" selection="multiple" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import SelectUserDialog from 'src/components/GqaSeleteUser/SelectUserDialog'
import { ref, computed, onMounted } from 'vue';
import { postAction } from 'src/api/manage';
import { useI18n } from 'vue-i18n';

const $q = useQuasar()
const { t } = useI18n()

const url = {
    list: 'plugin-vote/candidate-list',
    add: 'plugin-vote/candidate-add',
    delete: 'plugin-vote/candidate-delete',
}
const columns = computed(() => {
    return [
        { name: 'avatar', align: 'center', label: t('Avatar'), field: 'avatar' },
        { name: 'username', align: 'center', label: t('Username'), field: 'username' },
        { name: 'nickname', align: 'center', label: t('Nickname'), field: 'nickname' },
        { name: 'realName', align: 'center', label: t('RealName'), field: 'realName' },
        { name: 'voteType', align: 'center', label: '投票类型', field: 'voteType' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    showDateTime,
    dictOptions,
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
    queryParams.value.voteType = 'dy'
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    })
})

const selectUserDialog = ref(null)
const showAddUserForm = () => {
    selectUserDialog.value.show()
}

const handleAddUser = (event) => {
    const usernameList = []
    for (let i of event) {
        usernameList.push(i.username)
    }
    postAction(url.add, {
        candidate: usernameList,
        voteType: queryParams.value.voteType,
    }).then((res) => {
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: res.message,
            })
        }
        onRequest({
            pagination: pagination.value,
            queryParams: queryParams.value
        })
    })
}
</script>
