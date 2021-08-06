<template>
  <div style="margin-left: 50px; margin-top: 20px">
    <el-form
      ref="dataForm"
      :model="form"
      label-position="left"
      label-width="70px"
      style="width: 400px"
    >
      <el-form-item label="旧的密码" prop="oldPwd">
        <el-input
          v-model="form.oldPwd"
          placeholder="请输入旧的密码"
          show-password
          clearable
        />
      </el-form-item>
      <el-form-item label="新的密码" prop="newPwd">
        <el-input
          v-model="form.newPwd"
          placeholder="请输入新的密码"
          show-password
          clearable
        />
      </el-form-item>
      <el-form-item label="确认密码" prop="newPwd2">
        <el-input
          v-model="form.newPwd2"
          placeholder="请输入新的密码"
          show-password
          clearable
        />
      </el-form-item>
    </el-form>

    <el-button type="primary" @click="handle_form_newPwd()">
      确定修改
    </el-button>
  </div>
</template>

<script>
import { setToken } from "@/utils/auth";
import { HTTP_NEWPWD } from "@/network/newPwd.js";
import {
  HTTP_CALLBACK,
  HTTP_MSG_ERROR,
  HTTP_MSG_SUCCESS,
} from "@/network/utils.js";
export default {
  name: "newPwd",
  data() {
    return {
      form: {
        oldPwd: "",
        newPwd: "",
        newPwd2: "",
      },
    };
  },
  methods: {
    handle_form_newPwd() {
      //console.logog(this.form, this.form.newPwd2)
      if (this.form.oldPwd.length == 0 || this.form.newPwd.length == 0) {
        HTTP_MSG_ERROR("密码不可为空");
        return;
      } else if (this.form.newPwd != this.form.newPwd2) {
        HTTP_MSG_ERROR("两次输入的密码不一致");
        return;
      }
      HTTP_CALLBACK(
        HTTP_NEWPWD(this.form),
        (function (obj) {
          return function (res) {
            setToken("");
            window.top.location.href = "#/login";
          };
        })(this)
      );
    },
  },
};
</script>

<style>
</style>
