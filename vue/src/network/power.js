import {
  POST,
  GET
} from './request.js';

export function HTTP_GETALL(data) {
  return GET({
    url: "/v1/power/getAll",
    params: {
      page: data.page,
      limit: data.limit,
      aid: data.aid,
      state: data.state,
      uid: data.uid
    }
  })
}
export function HTTP_CHARGE(data) {
  return POST({
    url: "/v1/power/charge",
    data: {
      user: data.user,
      cdk: data.cdk,
      aid: data.aid,
    }
  })
}

export function HTTP_EDIT(data) {
  return POST({
    url: "/v1/power/edit",
    data: {
      pid: data.pid,
      state: data.state,
      day: data.day,
      point: data.point,
      ip: data.ip,
      code: data.code
    }
  })
}
