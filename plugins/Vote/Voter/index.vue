<template>
    <q-page padding>

        <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.voter" label="固定投票人" clearable />
            <q-select style="width: 20%" v-model="queryParams.voteType" :options="dictOptions.voteType" emit-value
                map-options label="投票类型" @update:model-value="handleSearchWithRatio" />
            <q-select style="width: 20%" v-if="queryParams.voteType === 'dy'" v-model="queryParams.voteRatio"
                :options="dictOptions.voteDyRatio" emit-value map-options label="投票权重"
                @update:model-value="handleSearch" />
            <q-select style="width: 20%" v-if="queryParams.voteType === 'gl'" v-model="queryParams.voteRatio"
                :options="dictOptions.voteGlRatio" emit-value map-options label="投票权重"
                @update:model-value="handleSearch" />
            <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
            <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <div class="q-gutter-xs" v-if="queryParams.voteType === 'dy'">
                    <q-btn v-for="(item, index) in dictOptions.voteDyRatio" :key="index" dense color="primary"
                        @click="showAddUserForm(item.dict_code, item.dict_label)">
                        新增
                        {{ item.dict_label }}({{ item.dict_ext_1 }}%)
                    </q-btn>
                </div>

                <div class="q-gutter-xs" v-if="queryParams.voteType === 'gl'">
                    <q-btn v-for="(item, index) in dictOptions.voteGlRatio" :key="index" dense color="primary"
                        @click="showAddUserForm(item.dict_code, item.dict_label)">
                        新增
                        {{ item.dict_label }}({{ item.dict_ext_1 }}%)
                    </q-btn>
                </div>

                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
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
                <q-td :props="props" v-if="queryParams.voteType === 'dy'">
                    <GqaDictShow dictName="voteDyRatio" :dictCode="props.row.voteRatio" withExt1 ext1="%" />
                </q-td>
                <q-td :props="props" v-if="queryParams.voteType === 'gl'">
                    <GqaDictShow dictName="voteGlRatio" :dictCode="props.row.voteRatio" withExt1 ext1="%" />
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

        <SelectUserDialog ref="selectUserDialog" @handleSelectUser="handleAddUser" selection="multiple"
            :title="selectRatioLabel" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import SelectUserDialog from 'src/components/GqaSeleteUser/SelectUserDialog'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'plugin-vote/voter-list',
    add: 'plugin-vote/voter-add',
    delete: 'plugin-vote/voter-delete',
}
const columns = computed(() => {
    return [
        { name: 'avatar', align: 'center', label: t('Avatar'), field: 'avatar' },
        { name: 'username', align: 'center', label: t('Username'), field: 'username' },
        { name: 'nickname', align: 'center', label: t('Nickname'), field: 'nickname' },
        { name: 'realName', align: 'center', label: t('RealName'), field: 'realName' },
        { name: 'voteType', align: 'center', label: '投票类型', field: 'voteType' },
        { name: 'voteRatio', align: 'center', label: '投票权重', field: 'voteRatio' },
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
    getTableData,
    handleSearch,
    // resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    queryParams.value.voteType = 'dy'
    queryParams.value.voteRatio = ''
    getTableData()
})
const selectRatio = ref('')
const selectRatioLabel = ref('')

const resetSearch = () => {
    queryParams.value = {
        voteType: 'dy',
        voteRatio: '',
    }
    getTableData()
}
const handleSearchWithRatio = () => {
    queryParams.value.voteRatio = ''
    getTableData()
}
const selectUserDialog = ref(null)
const showAddUserForm = (dictCode, dictLabel) => {
    selectRatio.value = dictCode
    selectRatioLabel.value = dictLabel
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
        voteRatio: selectRatio.value,
    }).then((res) => {
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: res.message,
            })
        }
        getTableData()
    })
}
</script>
