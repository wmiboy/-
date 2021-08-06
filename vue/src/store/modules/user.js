import {
  getToken,
  setToken,
  removeToken,
  setRouter,
} from '@/utils/auth'
import router, {
  resetRouter
} from '@/router'

import {stringify} from '../../permission.js'
const state = {
  token: '111',
  name: '',
  avatar: '',
  introduction: '',
  roles: [],
  isRouter: false,
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_INTRODUCTION: (state, introduction) => {
    state.introduction = introduction
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_isRouter: (state, isRouter) => {
    state.isRouter = isRouter
  }
}

const actions = {
  setUser({ commit }, user) {
    commit('SET_NAME', user.name)
    commit('SET_ROLES', user.roles)
    commit('SET_TOKEN', user.token)
    commit('SET_AVATAR', user.avatar)
    setToken(user.token)
    setRouter(stringify({isRouter:false,roles:user.roles}))
  },
  setisRouter({ commit }, isRouter) {
    commit('SET_isRouter', isRouter)
  },
  resetToken({ commit }) {
    return new Promise(resolve => {
      commit('SET_TOKEN', '')
      commit('SET_ROLES', [])
      commit('SET_NAME', '')
      commit('SET_AVATAR', '')
      commit('SET_isRouter', false)
      removeToken()
      resolve()
      setRouter(stringify({isRouter:false,roles:[]}))
    })
  },
  // dynamically modify permissions
  async changeRoles({ commit, dispatch }, role) {
    const token = role + '-token'
    commit('SET_TOKEN', token)
    setToken(token)

    const {
      roles
    } = await dispatch('getInfo')

    resetRouter()

    // generate accessible routes map based on roles
    const accessRoutes = await dispatch('permission/generateRoutes', roles, {
      root: true
    })
    // dynamically add accessible routes
    router.addRoutes(accessRoutes)

    // reset visited views and cached views
    dispatch('tagsView/delAllViews', null, {
      root: true
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
