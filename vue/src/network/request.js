import Axios from 'axios';
import { getToken, setToken } from '@/utils/auth';
//const path = "http://127.0.0.1:52777";
const path = "";
export function POST(conf) {
  let http = Axios.create({
    method: "post",
    baseURL: path,
    timeout: 30000,
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    transformRequest: [function (data) {
      let ret = ''
      for (let it in data) {
        ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
      }
      return ret
    }]
  })

  return request(http, conf)


}

export function GET(conf) {
  const http = Axios.create({
    method: "get",
    baseURL: path,
    timeout: 30000
  })
  return request(http, conf)
}

function request(http, conf) {
  //请求拦截器

  http.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    //console.log(config)
    config.url = config.url + '?token=' + getToken()
    return config;
  },
    function (error) {
      // 对请求错误做些什么
      return Promise.reject(error);
    })


  //结果拦截器
  http.interceptors.response.use(function (config) {
    // 对响应数据做点什么
    //console.log(config);
    if (config.data.code == 600) {
      setToken('');
      window.top.location.href = '#/login';
    }
    return config;
  },
    function (error) {
      // 对响应错误做点什么
      return Promise.reject(error);
    })
  return http(conf)
}