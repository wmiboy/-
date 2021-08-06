<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input
        v-model="listQuery.user"
        placeholder="用户名"
        style="width: 200px; margin-right: 10px"
        class="filter-item"
        clearable
        @keyup.enter.native="handle_getAll"
      />
      <el-select
        v-model="listQuery.state"
        placeholder="用户状态"
        clearable
        class="filter-item"
        style="width: 120px; margin-right: 10px"
      >
        <el-option
          v-for="item in state"
          :key="item.key"
          :label="item.msg"
          :value="item.index"
        />
      </el-select>
      <el-button
        v-waves
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handle_getAll"
      >
        查询
      </el-button>
    </div>

    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      border
      fit
      stripe
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column label="用户ID" prop="uid" align="center" width="80">
        <template slot-scope="{ row }">
          <span>{{ row.uid }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="用户名"
        width="200px"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>{{ row.user }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="用户信息"
        min-width="150px"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>{{ row.msg }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="状态"
        class-name="status-col"
        align="center"
        width="95"
      >
        <template slot-scope="{ row }">
          <el-switch
            :value="row.state == 1"
            active-color="#13ce66"
            inactive-color="#ff4949"
            @change="handle_sw_edit_state(row)"
          >
          </el-switch>
        </template>
      </el-table-column>
      <el-table-column label="时间" width="100" align="center">
        <template slot-scope="{ row }">
          <span>{{ row.time | parseTime("{y}-{m}-{d}") }}</span>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total > 0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="handle_getAll"
    />
  </div>
</template>

<script>
import waves from "@/directive/waves"; // waves directive
import { parseTime } from "@/utils";
import Pagination from "@/components/Pagination"; // secondary package based on el-pagination
import { HTTP_GETALL, HTTP_ADD, HTTP_EDIT } from "@/network/user.js";
import {
  HTTP_CALLBACK,
  HTTP_MSG_ERROR,
  HTTP_MSG_SUCCESS,
} from "@/network/utils.js";

export default {
  name: "user",
  components: { Pagination },
  directives: { waves },
  data() {
    return {
      page: 1,
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 30,
        user: "",
        state: "",
      },
      temp: {
        state: 0,
        uid: 0,
      },
      state: [
        { index: 1, msg: "正常使用", key: "state_ok" },
        { index: 2, msg: "禁止使用", key: "state_no" },
      ],
      dialogPvVisible: false,
      pvData: [],
      rules: {},
      downloadLoading: false,
    };
  },
  created() {
    this.handle_getAll();
  },
  methods: {
    handle_sw_edit_state(row) {
      var state = row.state == 1 ? 2 : 1;
      this.temp.state = state;
      this.temp.uid = row.uid;
      HTTP_CALLBACK(
        HTTP_EDIT(this.temp),
        (function (obj, state) {
          return function (res) {
            obj.state = state;
          };
        })(row, state)
      );
    },
    handle_getAll() {
      if (this.listQuery.page > 1 && this.list.length < this.listQuery.limit) {
        HTTP_MSG_ERROR("别点了已经没有了");
        return;
      }
      this.page = this.listQuery.page;
      this.listLoading = true;
      HTTP_CALLBACK(
        HTTP_GETALL(this.listQuery),
        (function (obj) {
          return function (res) {
            obj.list = res.data;
            obj.total = res.count;
            obj.listLoading = false;
          };
        })(this),
        (function (obj) {
          return function (res) {
            obj.listLoading = false;
          };
        })(this)
      );
    },
    handle_query() {
      //查询事件
      this.listQuery.page = 1;
      this.handle_getAll();
    },
  },
};
</script>
