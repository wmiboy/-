import {
  POST,
  GET
} from './request.js';

export function HTTP_GETALL(data) {
  return GET({
    url: "/v1/apply/getAll",
    params: {
      page: data.page,
      limit: data.limit,
      aid: data.aid,
      state: data.state,
      name: data.name
    }
  })
}
export function HTTP_ADD(data) {
  return POST({
    url: "/v1/apply/add",
    data: {
      name: data.name,
      mold: data.mold,
      bin: data.bin,
      login: data.login,
      msg: data.msg
    }
  })
}

export function HTTP_EDIT(data) {
  return POST({
    url: "/v1/apply/edit",
    data: {
      aid: data.aid,
      name: data.name,
      mold: data.mold,
      bin: data.bin,
      login: data.login,
      msg: data.msg,
      state: data.state
    }
  })
}
