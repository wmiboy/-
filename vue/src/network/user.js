import {
  POST,
  GET
} from './request.js';

export function HTTP_GETALL(data) {
  return GET({
    url: "/v1/user/getall",
    params: {
      page: data.page,
      limit: data.limit,
      uid: data.uid,
      state: data.state,
      user: data.user
    }
  })
}
export function HTTP_EDIT(data) {
  return POST({
    url: "/v1/user/edit",
    data: {
      uid: data.uid,
      state: data.state
    }
  })
}

export function HTTP_GET(data) {
  return POST({
    url: "/v1/power/getAll",
    data: {
      uid: data.uid
    }
  })
}

