<template>
  <div style="margin: 20px">
    <el-row :gutter="10">
      <el-col :span="12">
        <div class="grid-content bg-purple">
          <el-card class="box-card" shadow="hover">
            <div slot="header" class="clearfix">
              <span>文档说明</span>
            </div>
            <h4>开始使用</h4>
            1.解压压缩包<br /><br />
            2.导入ruan_jian_guan_li.sql,添加新的数据库<br /><br />
            3.修改配置文件,详见下方配置文件说明<br /><br />
            4.运行web.exe启动<br /><br />
            5.前端文件在dist目录下,请配置NG访问,并转发/v1所有请求至web端口<br /><br />
            <h4><font color="#FF0000">开箱即用法 </font></h4>

            1.解压所有文件<br />
            2.打开配置文件,修改ismysql配置为2<br />
            3.按照ruan_jian_guan_li.sql的结构,配置一下sqlite3数据库.<br />
            4.将数据库文件名更名为data.db 与web.exe 放在同一目录<br />
            注:开箱即用法后台地址为http://localhost:9527/admin.html<br />

            <h4>入口说明</h4>
            1.后台入口: http://localhost:9527/<br /><br />
            <font color="#FF0000"
              >后台默认账号密码为 admin 123456,请登录后自行修改</font
            >
            <br /><br />
            2.API接口说明:http://127.0.0.1:52777/api/swagger/index.html/<br /><br />
            这是后台所有的API接口,可进行二次开发
            <br /><br />

            <h4>业务流程</h4>
            1.通过[登录接口]登录获取token<br /><br />
            2.轮训[心跳接口]提交登录token<br /><br />
            3.通过[心跳接口]返回数据,判断用户状态/用户到期时间/用户剩余点数/应用状态...来决定应用继续使用还是关闭及其他操作<br /><br />
            4.通过[心跳接口]返回数据,获取应用对应的远程数据,建议应用的关键数据均放在此处联网获取<br /><br />
            5.若收费模式为[扣点]可通过[扣点接口]进行扣点操作<br /><br />
            <font color="#FF0000">
              注1:登录/心跳/扣点均做了通讯加密,可保障数据安全<br /><br />
              注2:所有接口均为http接口,可对接PC端/APP端/小程序等等多种需要网络验证的程序
            </font>
            <br /><br />
            <h4>响应代码说明</h4>
            {"code":200,"msg":"扣除成功"}<br />
            {"code":400,msg:"账户名或密码错误"}<br />
            code-200[成功]<br />
            code-300[成功但是响应数据已加密需要解密,详见通讯加密]<br />
            code-400[失败]<br />
            code-600[token过期了]<br />
            code-800[加密异常]<br />
            ----请求异常-1001:RSA解密失败<br />
            ----请求异常-1002:AES解密失败<br />
            ----请求异常-1005:数据为空<br />
            ----请求异常-1006:签名验证不通过<br />
            ----请求异常-1007:请求时间与服务器时间间隔大于2分钟
            <br /><br />
            <h4>配置文件说明</h4>
            目录下conf.yaml文件,需按照yaml文件格式修改配置,修改后需重启后台使配置生效<br />
            ismysql-默认即可<br />
            port-开放端口<br />

            mysql.host-mysql地址<br />
            mysql.name-数据库名<br />
            mysql.user-数据库用户名<br />
            mysql.pwd-数据库密码<br />
            mysql.maxconn-最大空闲连接数<br />
            mysql.maxopen-最大连接数<br /><br />

            encrypt.salt-签名加密的盐值<br />
            encrypt.privateKey-RSA私钥<br />
            encrypt.publicKey-RSA公钥
            <br /><br />
            <h4>接口相关</h4>
            下方接口带★均需加密提交,详见通讯加密<br /><br />
            <el-collapse>
              <el-collapse-item title="★[POST]应用查询接口" name="7">
                http://127.0.0.1:52777/v1/apply/get<br />
                aid=应用ID<br />
                {"aid":10002,"name":"xxxxx助手","mold":1,"login":1,"bin":1,"msg":"66666","state":1,"time":1627215600}<br />
                {"code":400,"msg":"收费软件不可直接获取信息"}
                {"code":400,"msg":"软件不存在"}
              </el-collapse-item>
              <el-collapse-item title="★[POST]登录接口" name="1">
                http://192.168.1.7:52777/v1/power/login<br />
                user=账号&pwd=密码&aid=应用ID&dev=设备特征(机器码)<br />
                {"token":"登录token"}<br /><br />
                常见错误返回<br />
                {"code":400,msg:"参数异常"}<br />
                {"code":400,msg:"账户名或密码错误"}<br />
                {"code":400,msg:"账户已被禁用"}<br />
                {"code":400,msg:"当前设备与登录设备不一致"}<br />
                {"code":400,msg:"应用不存在"}<br />
                {"code":400,msg:"应用暂不可用"}<br />
                {"code":400,msg:"应用无需登录"}<br />
                {"code":400,msg:"账户未充值,请充值后使用"}<br />
                {"code":400,msg:"账户已被禁止使用此应用"}<br />
                {"code":400,msg:"账户已到期,请充值后使用"}<br />
                {"code":400,msg:"账户剩余点数不足,请充值后使用"}<br />
                {"code":400,msg:"非绑定设备无法登录"}<br />
                {"code":400,msg:"非绑定IP无法登录"}<br />
              </el-collapse-item>
              <el-collapse-item title="★[POST]心跳接口" name="2">
                http://192.168.1.7:52777/v1/power/check?token=<br />
                dev=设备特征(机器码)<br />
                {"a":{"aid":应用ID,"name":"应用名称","mold":收费模式,"login":登录模式,"bin":绑定模式,"msg":"远程数据","state":应用状态,"time":11},"p":{"pid":6,"uid":用户id,"aid":应用id,"kid":卡密id,"day":过期时间,"point":剩余点数,"ip":"绑定IP","code":"绑定设备","state":权限状态,"time":1625560772}}
                <br /><br />
                参数说明<br />
                mold:1-扣点 2-扣时<br />
                login:1-免登陆 2-单点登录 3-多点登录<br />
                bin:1-不绑定 2-绑定机器 3-绑定IP<br />
                state:1-正常 2-禁用<br />
                day:到期时间的10位时间戳<br /><br />
                常见错误返回<br />
                {"code":400,msg:"当前设备与登录设备不一致"}<br />
                {"code":400,msg:"应用不存在"}<br />
                {"code":400,msg:"应用暂不可用"}<br />
                {"code":400,msg:"应用无需登录"}<br />
                {"code":400,msg:"账户未充值,请充值后使用"}<br />
                {"code":400,msg:"账户已被禁止使用此应用"}<br />
                {"code":400,msg:"账户已到期,请充值后使用"}<br />
                {"code":400,msg:"账户剩余点数不足,请充值后使用"}<br />
                {"code":400,msg:"非绑定设备无法登录"}<br />
                {"code":400,msg:"非绑定IP无法登录"}<br />
              </el-collapse-item>

              <el-collapse-item title="★[POST]扣点接口" name="3">
                http://192.168.1.7:52777/v1/power/sub?token=<br />
                point=欲扣除的点数,不可为负数<br />
                {"code":200,"msg":"扣除成功"}
                <br /><br />常见错误返回<br />
                {"code":400,"msg":"扣除失败,剩余点数不足"}
              </el-collapse-item>
              <el-collapse-item title="[POST]注册接口" name="4">
                http://192.168.1.7:52777/v1/user/reg<br />
                user=账号&pwd=密码
              </el-collapse-item>
              <el-collapse-item title="[POST]改密接口" name="5">
                http://192.168.1.7:52777/v1/user/newPwd?token=<br />
                oldPwd=旧密码&newPwd=新密码
              </el-collapse-item>
              <el-collapse-item title="[POST]充值接口" name="6">
                http://192.168.1.7:52777/v1/power/charge<br />
                user=充值的账号&cdk=充值的卡密&aid=充值的应用ID
              </el-collapse-item>

              <el-collapse-item title="[GET]获取服务器时间" name="8">
                http://192.168.1.7:52777/v1/time<br />
                1625565258
              </el-collapse-item>
              <el-collapse-item title="通讯加密" name="9">
                请求报文:k=&d=&t=&s=
                <font color="#FF0000">参数均需进行URL编码-UTF8</font
                ><br /><br />

                [RSA_publicKey]RSA公钥,需于服务器中的RSA私钥匹配--欲修改见配置文件说明<br />
                [AES_KEY]AES公钥,随机16位字符即可,需保存用于解密本次通讯中的响应数据<br />
                [AES_IV]AES向量IV,固定为r6rS1jOno6zSDonx<br />
                [SIGN_TIME]通过接口获取的服务器时间<br />
                [SIGN_SALT]签名效验盐值--欲修改见配置文件说明<br />
                [Data]通讯的明文数据<br />
                <br />
                k=RSA_PCKS1(RSA_publicKey,AES_KEY)<br />d=AES_CBC_PCKS5(Data,AES_KEY,AES_IV)<br />t=SIGN_TIME<br />s=MD5(SIGN_TIME+SIGN_SALT)<br /><br />
                解密只需解密响应数据中的data字段即可<br />
                状态码不为300时表示该响应数据无需解密,详见响应代码说明<br />
              </el-collapse-item>
            </el-collapse>
          </el-card>
        </div>
      </el-col>
      <el-col :span="12"
        ><div class="grid-content bg-purple-light">
          <div class="grid-content bg-purple">
            <el-card class="box-card" shadow="hover">
              <div slot="header" class="clearfix">
                <span>更新日志</span>
              </div>
              <el-timeline>
                <el-timeline-item
                  v-for="(activity, index) in activities"
                  :key="index"
                  :icon="activity.icon"
                  :type="activity.type"
                  :color="activity.color"
                  :size="activity.size"
                  :timestamp="activity.timestamp"
                >
                  <el-card class="box-card" shadow="hover">
                    <div slot="header" class="clearfix">
                      <span>
                        <b>{{ activity.title }}</b></span
                      >
                    </div>
                    <div v-html="activity.content"></div>
                  </el-card>
                </el-timeline-item>
              </el-timeline>
            </el-card>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      activities: [
        {
          title: "V1.1 小改动",
          content: "1.增加用户信息模块,可简易存储需要联网保存的用户信息<br/><br/>2.开放API以供二此开发<br/><br/>3.增加开箱即用用法<br/><br/>4.优化代码<br/><br/>5.优化数据库结构<br/><br/>6.近期打算添加卡密登录/WebSocket对接",
          timestamp: "2021-07-07",
          size: "large",
          color: "#F56C6C",
        },
        {
          title: "V1.0 对接文档整理完毕",
          content: "对接步骤/对接接口/配置文件...",
          timestamp: "2021-07-07",
          size: "large",
          color: "#409EFF",
        },
        {
          title: "V1.0 前端开发完成",
          content:
            "1.完成WEB端后台管理开发<br/><br/>2.完成PC端对接例程<br/><br/>3.增加通讯加密防止数据被截获/篡改",
          timestamp: "2021-07-06",
          size: "large",
          color: "#409EFF",
        },
        {
          title: "V1.0 后台开发完成",
          content:
            "1.根据初版重写了一份正式版代码<br/><br/>2.统一了编码规范,优化精简代码<br/><br/>3.优化了部分业务逻辑",
          timestamp: "2021-07-01",
          color: "#409EFF",
        },
        {
          title: "V1.0 初版开发完毕",
          content:
            "1.摸鱼半个月,大体框架完成:应用管理/用户管理/卡密管理/权限管理<br/><br/>2.基础功能完成:用户注册/登录/充值/扣点/效验...<br/><br/>3.应用支持多种登录策略:免登陆/单点登录/多点登陆<br/><br/>4.应用支持多种绑定模式:不绑定/绑定机器码/绑定IP<br/><br/>5.应用支持多种收费方式:扣点消费/扣时消费<br/><br/>6.支持单账号多软件授权,一个账号可以登录所有软件,权限分开管理.精细化管理,再也不用一个软件注册一个账号了",
          timestamp: "2021-06-25",
          size: "large",
        },
      ],
    };
  },
};
</script>
<style>
</style>