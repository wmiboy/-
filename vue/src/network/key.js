import {
  POST,
  GET
} from './request.js';

export function HTTP_GETALL(data) {
  return GET({
    url: "/v1/key/getAll",
    params: {
      page: data.page,
      limit: data.limit,
      aid: data.aid,
      state: data.state,
      cdk: data.cdk
    }
  })
}
export function HTTP_EDIT(data) {
  return POST({
    url: "/v1/key/edit",
    data: {
      kid: data.kid,
      state: data.state
    }
  })
}

export function HTTP_CREAD(data) {  
  return POST({
    url: "/v1/key/cread",
    data: {
      num: data.num,
      mold: data.mold,
      point: data.point,
      day: data.day,
      aid: data.aid
    }
  })
}