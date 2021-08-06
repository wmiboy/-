import { MD5 } from './md5.js';
import {
  POST
} from './request.js';

export function HTTP_LOGIN(data) {
  return POST({
    url: "/v1/user/login",
    data: {
      user: data.user,
      pwd: MD5(data.pwd)
    }
  })
}
