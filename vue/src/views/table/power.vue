<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input
        v-model="listQuery.aid"
        placeholder="应用ID"
        style="width: 200px; margin-right: 10px"
        class="filter-item"
        clearable
        @keyup.enter.native="handle_query"
      />
      <el-input
        v-model="listQuery.uid"
        placeholder="用户ID"
        style="width: 200px; margin-right: 10px"
        class="filter-item"
        clearable
        @keyup.enter.native="handle_query"
      />
      <el-select
        v-model="listQuery.state"
        placeholder="使用状态"
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
      <el-table-column label="id" prop="pid" align="center" width="80">
        <template slot-scope="{ row }">
          <span>{{ row.pid }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="用户ID"
        width="80px"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>{{ row.uid }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="应用ID"
        width="70px"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>{{ row.aid }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="卡密ID"
        width="70px"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>{{ row.kid }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="到期时间"
        width="120px"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>{{ row.day | parseTime("{m}-{d} {h}:{i}") }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="剩余点数"
        width="80px"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>{{ row.point }}</span>
        </template>
      </el-table-column>

      <el-table-column
        label="绑定IP"
        width="150px"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>{{ row.ip }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="绑定设备"
        align="center"
        min-width="150px"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>{{ row.code }}</span>
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
      <el-table-column label="最近使用" width="100" align="center">
        <template slot-scope="{ row }">
          <span>{{ row.time | parseTime("{m}-{d} {h}:{i}") }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="操作"
        align="center"
        width="120px"
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
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left: 50px"
      >
        <el-form-item label="到期时间" prop="day">
          <el-date-picker
            v-model="time"
            type="datetime"
            value-format="timestamp"
            format="yyyy-MM-dd HH:mm"
            placeholder="选择到期时间"
          >
          </el-date-picker>
        </el-form-item>
        <el-form-item label="可用点数" prop="point">
          <el-input v-model="temp.point" placeholder="请输入欲修改的点数" />
        </el-form-item>
        <el-form-item label="绑定IP" prop="ip">
          <el-input v-model="temp.ip" placeholder="请输入欲修改的IP" />
        </el-form-item>
        <el-form-item label="绑定设备" prop="code">
          <el-input v-model="temp.code" placeholder="请输入欲修改的设备" />
        </el-form-item>
      </el-form>
      <div style="margin-left: 50px">
        <el-button
          type="primary"
          @click="dialogStatus === 'create' ? handle_from_add() : handle_from_edit()"
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
import { HTTP_GETALL, HTTP_ADD, HTTP_EDIT } from "@/network/power.js";
import {
  HTTP_CALLBACK,
  HTTP_MSG_ERROR,
  HTTP_MSG_SUCCESS,
} from "@/network/utils.js";

export default {
  name: "power",
  components: { Pagination },
  directives: { waves },
  data() {
    return {
      time: 0,
      page: 1,
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 30,
        aid: "",
        uid: "",
        state: "",
      },
      state: [
        { index: 1, msg: "正常", key: "state_ok" },
        { index: 2, msg: "禁用", key: "state_no" },
      ],
      temp: {
        pid: "",
        state: "",
        day: "",
        point: "",
        ip: "",
        code: "",
      },
      dialogFormVisible: false,
      dialogStatus: "",
      textMap: {
        update: "编辑",
        create: "添加",
      },
    };
  },
  created() {
    this.handle_getAll();
  },
  methods: {
    resetTemp() {
      this.temp = {
        pid: "",
        state: "",
        day: "",
        point: "",
        ip: "",
        code: "",
      };
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
      this.time = row.day * 1000;
      this.temp = Object.assign({}, row); // copy obj
      this.dialogStatus = "update";
      this.dialogFormVisible = true;
      this.$nextTick(() => {
        this.$refs["dataForm"].clearValidate();
      });
    },
    handle_from_edit() {
      this.temp.day = this.time / 1000;
      HTTP_CALLBACK(
        HTTP_EDIT(this.temp),
        (function (obj) {
          return function (res) {
            const index = obj.list.findIndex((v) => v.pid === obj.temp.pid);
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
      this.temp.pid = row.pid;
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
