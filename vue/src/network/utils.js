import { Message } from "element-ui";
export function HTTP_MSG_SUCCESS(msg) {
    Message.success({ message: msg, duration: 1000, offset: 200 })
}
export function HTTP_MSG_ERROR(msg) {
    Message.error({ message: msg, duration: 1000, offset: 200 })
}
export function HTTP_CALLBACK(promise, fun_success, fun_error) {
    /*
    回调函数格式
     var success = function (obj) {
        return function (res) {
        };
      }; 
    */
    promise.then((response) => {
        let res = response.data;
        if (res.code == 200) {
            HTTP_MSG_SUCCESS(res.msg);
            if (fun_success != null) { fun_success(res); }
        } else {
            HTTP_MSG_ERROR(res.msg);
            if (fun_error != null) { fun_error(res); }
        }
    }).catch((error) => {
        console.log(error);
        HTTP_MSG_ERROR("网络异常");
        return false;
    });
}
