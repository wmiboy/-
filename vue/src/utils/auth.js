import Cookies from 'js-cookie'
import store from '../store'

const TokenKey = 'RuanJianGunLi-Token'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

export function setUser(data) {
  store.dispatch('user/setUser', data)
}

export function setRouter(data) {
  return Cookies.set("Router", data)
}

export function getRouter(data) {
  return Cookies.get("Router")
}

