<template>
  <div>
    <a-table :columns="columns" :data-source="dataSource" :loading="loading" :pagination="false">
      <template #bodyCell="{column, record}">
        <template v-if="column.key === 'action'">
                    <span>
                        <a class="mr-2 font-medium" @click="handleTopicDetail(record.v)">Detail</a>
                        <a class="text-red-400 font-medium" @click="deleteTopic(record.k)">Delete</a>
                    </span>
        </template>
      </template>
    </a-table>
    <div class="flex justify-center mt-2 text-lg pager" v-if="canPrev || canNext">
      <a-button type="text" :disabled="!canPrev" @click="handlePrevPage">Previous</a-button>
      <a-button type="text" :disabled="!canNext" @click="handleNextPage">Next</a-button>
    </div>

    <DialogTopicDetail ref="topicDetail"/>
  </div>
</template>
<script setup lang="ts">
import {IListAllKVReq, listKVPairs,deleteTopic} from '/@/api/api_dtm'
import {computed, ref} from 'vue-demi'
import {usePagination} from 'vue-request'
import DialogTopicDetail from './_Components/DialogTopicDetail.vue';

const columns = [
  {
    title: 'Name',
    dataIndex: 'k',
    key: 'name'
  }, {
    title: 'Subscribers',
    dataIndex: 'v',
    key: 'subscribers'
  }, {
    title: 'Version',
    dataIndex: 'version',
    key: 'version'
  }, {
    title: 'Action',
    key: 'action'
  }
]

const pages = ref([''])
const curPage = ref(1)

const canPrev = computed(() => {
  return curPage.value > 1
})

const canNext = computed(() => {
  return data.value?.data.next_position !== ""
})

type Data = {
  kv: {
    k: string
    v: string
    version: number
  }[]
  next_position: string
}

const queryData = (params: IListAllKVReq) => {
  return listKVPairs<Data>(params)
}

const {data, run, current, loading, pageSize} = usePagination(queryData, {
  defaultParams: [
    {
      cat: "topics",
      limit: 10,
    }
  ],
  pagination: {
    pageSizeKey: 'limit'
  }
})

const dataSource = computed(() => data.value?.data.kv || [])

const handlePrevPage = () => {
  curPage.value -= 1;
  const params = {
    cat: "topics",
    limit: pageSize.value,
    position: pages.value[curPage.value] as string
  }
  run(params)
}

const handleNextPage = () => {
  curPage.value += 1;
  pages.value[curPage.value] = data.value?.data.next_position as string

  run({
    cat: "topics",
    position: data.value?.data.next_position,
    limit: pageSize.value,
  })
}

const topicDetail = ref<null | { open: (subscribers: string) => null }>(null)
const handleTopicDetail = (subscribers: string) => {
  topicDetail.value?.open(subscribers)
}

</script>

<style lang="postcss" scoped>
::v-deep .ant-pagination-item {
  display: none;
}

.pager .ant-btn-text {
  font-weight: 500;
}

.pager .ant-btn {
  padding: 6px;
}
</style>
