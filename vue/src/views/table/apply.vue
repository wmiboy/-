<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input
        v-model="listQuery.name"
        placeholder="应用名称"
        style="width: 200px; margin-right: 10px"
        class="filter-item"
        clearable
        @keyup.enter.native="handle_query"
      />
      <el-select
        v-model="listQuery.state"
        placeholder="应用状态"
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
        @click="handle_query"
      >
        查询
      </el-button>
      <el-button
        class="filter-item"
        style="margin-left: 10px"
        type="primary"
        icon="el-icon-circle-plus-outline"
        @click="handle_but_add"
      >
        添加
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
      <el-table-column label="应用ID" prop="id" align="center" width="80">
        <template slot-scope="{ row }">
          <span>{{ row.aid }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="应用名称"
        width="150px"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>{{ row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="远程数据"
        min-width="150px"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>
            {{ row.msg }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label="绑定模式" width="95px" align="center">
        <template slot-scope="{ row }">
          <el-tag :type="get_bin(row.bin).tag">
            <span>{{ get_bin(row.bin).msg }}</span>
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="收费模式" width="95px" align="center">
        <template slot-scope="{ row }">
          <el-tag :type="get_mold(row.mold).tag" effect="dark">
            <span>{{ get_mold(row.mold).msg }}</span>
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="登录模式" width="95px" align="center">
        <template slot-scope="{ row }">
          <el-tag :type="get_login(row.login).tag">
            <span>{{ get_login(row.login).msg }}</span>
          </el-tag>
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
      <el-table-column
        label="操作"
        align="center"
        width="100"
        class-name="small-padding fixed-width"
        fixed="right"
      >
        <template slot-scope="{ row }">
          <el-button type="primary" size="mini" @click="handle_but_edit(row)">
            编辑
          </el-button>
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

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left: 50px"
      >
        <el-form-item label="应用名称" prop="name">
          <el-input v-model="temp.name" placeholder="请输入应用名称" />
        </el-form-item>
        <el-form-item label="登录策略" prop="state">
          <el-select
            v-model="temp.login"
            class="filter-item"
            placeholder="请选择登录策略"
          >
            <el-option
              v-for="item in login"
              :key="item.key"
              :label="item.msg"
              :value="item.index"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="绑定策略" prop="bin">
          <el-select
            v-model="temp.bin"
            class="filter-item"
            placeholder="请选择绑定策略"
          >
            <el-option
              v-for="item in bin"
              :key="item.key"
              :label="item.msg"
              :value="item.index"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="收费策略" prop="mold">
          <el-select
            v-model="temp.mold"
            class="filter-item"
            placeholder="请选择收费策略"
          >
            <el-option
              v-for="item in mold"
              :key="item.key"
              :label="item.msg"
              :value="item.index"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="应用状态" prop="state">
          <el-select
            v-model="temp.state"
            class="filter-item"
            placeholder="请选择应用状态"
          >
            <el-option
              v-for="item in state"
              :key="item.key"
              :label="item.msg"
              :value="item.index"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="远程数据">
          <el-input
            v-model="temp.msg"
            :autosize="{ minRows: 3, maxRows: 5 }"
            type="textarea"
            placeholder="可将关键信息放在此处,由应用联网动态获取"
          />
        </el-form-item>
      </el-form>
      <div style="margin-left: 50px">
        <el-button
          type="primary"
          @click="
            dialogStatus === 'create' ? handle_from_add() : handle_from_edit()
          "
        >
          确定
        </el-button>
        <el-button @click="dialogFormVisible = false"> 取消 </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import waves from "@/directive/waves"; // waves directive
import { parseTime } from "@/utils";
import Pagination from "@/components/Pagination"; // secondary package based on el-pagination
import { HTTP_GETALL, HTTP_ADD, HTTP_EDIT } from "@/network/apply.js";
import {
  HTTP_CALLBACK,
  HTTP_MSG_ERROR,
  HTTP_MSG_SUCCESS,
} from "@/network/utils.js";
const enum_state = [
  { msg: "正常", tag: "success" },
  { msg: "停用", tag: "danger" },
];
const enum_login = [
  { msg: "无需登录", tag: "info" },
  { msg: "单点登录", tag: "success" },
  { msg: "多点登录", tag: "success" },
];
const enum_bin = [
  { msg: "不绑定", tag: "info" },
  { msg: "机器码", tag: "" },
  { msg: "服务器", tag: "" },
];
const enum_mold = [
  { msg: "扣时", tag: "" },
  { msg: "扣点", tag: "success" },
];
export default {
  name: "apply",
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
        name: "",
        state: "",
      },
      state: [
        { index: 1, msg: "正常", key: "state_ok" },
        { index: 2, msg: "暂停", key: "state_no" },
      ],
      login: [
        { index: 1, msg: "无需登录", key: "login_no" },
        { index: 2, msg: "单点登录", key: "login_dan" },
        { index: 3, msg: "多点登录", key: "loign_duo" },
      ],
      bin: [
        { index: 1, msg: "不绑定", key: "bin_no" },
        { index: 2, msg: "绑定机器", key: "bin_dev" },
        { index: 3, msg: "绑定IP", key: "bin_ip" },
      ],
      mold: [
        { index: 1, msg: "扣时", key: "mold_day" },
        { index: 2, msg: "扣点", key: "mold_point" },
      ],
      statusOptions: ["published", "draft", "deleted"],
      showReviewer: false,
      temp: {
        aid: "",
        name: "",
        state: 1,
        login: 1,
        bin: 1,
        mold: 1,
        msg: "",
      },
      dialogFormVisible: false,
      dialogStatus: "",
      textMap: {
        update: "编辑",
        create: "添加",
      },
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
    get_login(index) {
      if (index <= 0) {
        return "info";
      }
      return enum_login[index - 1];
    },
    get_state(index) {
      if (index <= 0) {
        return "info";
      }
      return enum_state[index - 1];
    },
    get_bin(index) {
      if (index <= 0) {
        return "info";
      }
      return enum_bin[index - 1];
    },
    get_mold(index) {
      if (index <= 0) {
        return "info";
      }
      return enum_mold[index - 1];
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
    resetTemp() {
      this.temp = {
        aid: "",
        name: "",
        state: 1,
        login: 1,
        bin: 1,
        mold: 1,
        msg: "",
      };
    },
    handle_but_add() {
      //添加事件
      this.resetTemp();
      this.dialogStatus = "create";
      this.dialogFormVisible = true;
      this.$nextTick(() => {
        this.$refs["dataForm"].clearValidate();
      });
    },
    handle_from_add() {
      HTTP_CALLBACK(
        HTTP_ADD(this.temp),
        (function (obj) {
          return function (res) {
            obj.dialogFormVisible = false;
          };
        })(this)
      );
    },
    handle_but_edit(row) {
      //修改事件
      this.temp = Object.assign({}, row); // copy obj
      this.dialogStatus = "update";
      this.dialogFormVisible = true;
      this.$nextTick(() => {
        this.$refs["dataForm"].clearValidate();
      });
    },
    handle_from_edit() {
      HTTP_CALLBACK(
        HTTP_EDIT(this.temp),
        (function (obj) {
          return function (res) {
            const index = obj.list.findIndex((v) => v.aid === obj.temp.aid);
            obj.list.splice(index, 1, obj.temp);
            obj.dialogFormVisible = false;
          };
        })(this)
      );
    },
    handle_sw_edit_state(row) {
      this.resetTemp();
      var state = row.state == 1 ? 2 : 1;
      this.temp.state = state;
      this.temp.aid = row.aid;
      HTTP_CALLBACK(
        HTTP_EDIT(this.temp),
        (function (obj, state) {
          return function (res) {
            obj.state = state;
          };
        })(row, state)
      );
    },
  },
};
</script>
