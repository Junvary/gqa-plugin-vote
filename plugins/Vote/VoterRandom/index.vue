<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.voter" label="投票人" clearable />
            <q-input style="width: 20%" v-model="queryParams.voteVersion" label="投票版本" clearable />
            <q-select style="width: 20%" v-model="queryParams.voteType" :options="dictOptions.voteType" emit-value
                map-options label="投票类型" @update:model-value="handleSearch" />

            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn dense color="primary" @click="showRandomUser()">
                    随机选取新一批投票人(
                    <GqaDictShow dictName="voteType" :dictCode="queryParams.voteType" />
                    )
                </q-btn>
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-createdAt="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.createdAt) }}
                </q-td>
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

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="negative" @click="handleDelete(props.row)" :label="$t('Delete')" />
                    </div>
                </q-td>
            </template>
        </q-table>

        <RandomUserDialog ref="randomUserDialog" @raomdomOver="getTableData" />
    </q-page>
</template>


<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import RandomUserDialog from './modules/RandomUserDialog'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'plugin-vote/voter-random-list',
    add: 'plugin-vote/voter-random-add',
    delete: 'plugin-vote/voter-random-delete',
}
const columns = computed(() => {
    return [
        { name: 'voteType', align: 'center', label: '投票类型', field: 'voteType' },
        { name: 'voteVersion', align: 'center', label: '投票版本', field: 'voteVersion' },
        { name: 'memo', align: 'center', label: '投票说明', field: 'memo' },
        { name: 'avatar', align: 'center', label: t('Avatar'), field: 'avatar' },
        { name: 'username', align: 'center', label: t('Username'), field: 'username' },
        { name: 'nickname', align: 'center', label: t('Nickname'), field: 'nickname' },
        { name: 'realName', align: 'center', label: t('RealName'), field: 'realName' },
        { name: 'createdAt', align: 'center', label: '创建时间', field: 'createdAt' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
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

onMounted(() => {
    queryParams.value.voteType = 'dy'
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    })
})
const getTableData = () => {
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    })
}
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
        voter: usernameList,
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
const randomUserDialog = ref(null)
const showRandomUser = () => {
    randomUserDialog.value.show(queryParams.value.voteType)
}
</script>
