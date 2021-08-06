import router from './router'
import store from './store'
import {
  Message
} from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import {
  getToken,
  setRouter,
  getRouter
} from '@/utils/auth' // get token from cookie
import defaultSettings from '@/settings'

const { title } = defaultSettings
NProgress.configure({
  showSpinner: false
}) // NProgress Configuration

const whiteList = ['/login'] // no redirect whitelist

router.beforeEach(async (to, from, next) => {
  //console.log(to)
  // 开启顶部进度条
  NProgress.start()
  let meta = to.meta
  console.log(meta)

  // 设置网页标题

    document.title = meta.title + "_" + title


  //获取token
  const hasToken = getToken()

  //未登录状态
  if (hasToken == null || hasToken.length == 0) {
    //判断当前连接是否为登录连接
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      //非登录连接，跳转到登录连接
      next("/login")
      NProgress.done()
    }
    return
  }
  //登录状态
  //已登录但是访问登录网址，跳转至首页
  if (to.path === '/login') {
    next({ path: '/' })
    NProgress.done()
    return
  }
  let isRouter = JSON.parse(getRouter())
  //console.log("getroles", isRouter, getRouter())
  if (isRouter.isRouter) {
    next()
    return
  }



  try {
    //获取用户权限
    const roles = isRouter.roles

    const accessRoutes = await store.dispatch('permission/generateRoutes', roles)
    router.addRoutes(accessRoutes)
    isRouter.isRouter = true
    setRouter(JSON.stringify(isRouter))
    next({ ...to, replace: true })
  } catch (error) {
    //添加路由失败动态，删除token重新登录
    await store.dispatch('user/resetToken')
    Message.error(error || 'Has Error')
    next('/login')
    NProgress.done()
  }

})

router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})

export function stringify(data) {
  return JSON.stringify(data)
}
