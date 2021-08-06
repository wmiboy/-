import { MD5 } from './md5.js';
import {
  POST
} from './request.js';

export function HTTP_NEWPWD(data) {
  return POST({
    url: "/v1/user/newPwd",
    data: {
      oldPwd: MD5(data.oldPwd),
      newPwd: MD5(data.newPwd)
    }
  })
}
