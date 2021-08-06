<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input
        v-model="listQuery.aid"
        placeholder="应用id"
        style="width: 200px; margin-right: 10px"
        class="filter-item"
        clearable
        @keyup.enter.native="handle_query"
      />
      <el-input
        v-model="listQuery.cdk"
        placeholder="卡密"
        style="width: 200px; margin-right: 10px"
        class="filter-item"
        clearable
        @keyup.enter.native="handle_query"
      />
      <el-select
        v-model="listQuery.state"
        placeholder="卡密状态"
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
        创建
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
      <el-table-column label="卡密ID" prop="id" align="center" width="80">
        <template slot-scope="{ row }">
          <span>{{ row.kid }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="卡密"
        min-width="150px"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>{{ row.cdk }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="应用ID"
        width="75px"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>
            {{ row.aid }}
          </span>
        </template>
      </el-table-column>
      <el-table-column
        label="类型"
        width="75px"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <el-tag :type="get_mold(row.mold).tag" effect="dark">
            <span>{{ get_mold(row.mold).msg }}</span>
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column
        label="小时 | 点数"
        width="120px"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <el-tag type="">{{ row.day }}</el-tag>

          <el-tag type="success" style="margin-left: 5px">{{
            row.point
          }}</el-tag>
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
            inactive-color="#909399"
            @change="handle_sw_edit_state(row)"
          >
          </el-switch>
        </template>
      </el-table-column>
      <el-table-column label="使用时间" width="100" align="center">
        <template slot-scope="{ row }">
          <span>{{ row.use_time | parseTime("{y}-{m}-{d}") }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="使用人"
        width="75px"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="{ row }">
          <span>
            {{ row.uid }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" width="100" align="center">
        <template slot-scope="{ row }">
          <span>{{ row.cread_time | parseTime("{y}-{m}-{d}") }}</span>
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
        style="width: 80%; margin-left: 50px"
      >
        <el-form-item label="卡密数量" prop="num">
          <el-input
            v-model="temp.num"
            placeholder="请输入生成卡密的数量"
            clearable
          />
        </el-form-item>
        <el-form-item label="卡密类型" prop="mold">
          <el-select
            v-model="temp.mold"
            class="filter-item"
            placeholder="请选择卡密类型"
          >
            <el-option
              v-for="item in mold"
              :key="item.key"
              :label="item.msg"
              :value="item.index"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="应用ID" prop="aid">
          <el-input
            v-model="temp.aid"
            placeholder="请输入卡密所对应的应用ID,0为通用卡密"
            clearable
          />
        </el-form-item>
        <el-form-item label="可用点数" prop="point">
          <el-input
            v-model="temp.point"
            placeholder="请输入卡密对应的点数"
            clearable
          />
        </el-form-item>
        <el-form-item label="可用小时" prop="day">
          <el-input
            v-model="temp.day"
            placeholder="请输入卡密对应的小时数"
            clearable
          />
        </el-form-item>
        <el-form-item label="卡密列表">
          <el-input
            v-model="temp.msg"
            :autosize="{ minRows: 10, maxRows: 10 }"
            type="textarea"
            placeholder="生成的卡密将会显示在这里"
          />
        </el-form-item>
      </el-form>
      <div style="margin-left: 50px">
        <el-button
          type="primary"
          :disabled="isDisabledBut"
          @click="dialogStatus === 'create' ? handle_from_add() : updateData()"
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
import { HTTP_GETALL, HTTP_CREAD, HTTP_EDIT } from "@/network/key.js";
import {
  HTTP_CALLBACK,
  HTTP_MSG_ERROR,
  HTTP_MSG_SUCCESS,
} from "@/network/utils.js";
const enum_state = [
  { msg: "未用", tag: "success" },
  { msg: "已用", tag: "danger" },
];
const enum_mold = [
  { msg: "扣时", tag: "" },
  { msg: "扣点", tag: "success" },
];
export default {
  name: "key",
  components: { Pagination },
  directives: { waves },
  data() {
    return {
      isDisabledBut: false,
      page: 1,
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 30,
        state: "",
        aid: "",
        cdk: "",
      },
      state: [
        { index: 1, msg: "未用", key: "state_ok" },
        { index: 2, msg: "已用", key: "state_no" },
      ],
      mold: [
        { index: 1, msg: "扣时", key: "mold_day" },
        { index: 2, msg: "扣点", key: "mold_point" },
      ],
      showReviewer: false,
      temp: {
        num: 10,
        mold: 1,
        point: "",
        day: "",
        aid: "",
        msg: "",
        kid: "",
      },
      dialogFormVisible: false,
      dialogStatus: "",
      textMap: {
        update: "编辑",
        create: "创建",
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
    resetTemp() {
      this.temp = {
        num: 10,
        mold: 1,
        point: "",
        day: "",
        aid: "",
        msg: "",
        kid: "",
      };
    },
    get_state(index) {
      if (index <= 0) {
        return "info";
      }
      return enum_state[index - 1];
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

    handle_but_add() {
      //添加事件
      this.resetTemp();
      this.dialogStatus = "create";
      this.dialogFormVisible = true;
    },
    handle_from_add() {
      if (this.temp.mold == 1 && this.temp.day == 0) {
        HTTP_MSG_ERROR("请填写卡密对应的小时数");
        return;
      } else if (this.temp.mold == 2 && this.temp.point == 0) {
        HTTP_MSG_ERROR("请填写卡密对应的点数");
        return;
      }

      HTTP_CALLBACK(
        HTTP_CREAD(this.temp),
        (function (obj) {
          return function (res) {
            var msg = "";
            for (var i = 0; i < res.data.length; i++) {
              var tmp = res.data[i];
              msg = msg + tmp.cdk + "----" + tmp.aid + "----";
              if (tmp.mold == 1) {
                msg = msg + "扣时:" + tmp.day + "小时----";
              } else {
                msg = msg + "扣点:" + tmp.point + "点----";
              }
              msg =
                msg +
                parseTime(tmp.cread_time, "{y}-{m}-{d} {h}:{i}:{s}") +
                "\n";
              obj.temp.msg = msg;
            }
            //obj.dialogFormVisible = false;
          };
        })(this)
      );
    },

    updata(data, row) {
      edit(data)
        .then((response) => {
          let res = response.data;
          if (res.code == 200) {
            Message.success({ message: res.msg, duration: 1000, offset: 200 });
            row.state = data.state;
          } else {
            Message.error({ message: res.msg, duration: 1000, offset: 200 });
          }
        })
        .catch((error) => {
          //console.log(error);
          Message.error({ message: "网络异常", offset: 200 });
          return false;
        });
    },

    handle_sw_edit_state(row) {
      this.resetTemp();
      var state = row.state == 1 ? 2 : 1;
      this.temp.state = state;
      this.temp.kid = row.kid;
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
