import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'


/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    noCache: true                if set true, the page will no be cached(default is false)
    affix: true                  if set true, the tag will affix in the tags-view
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
  constantRoutes
  无权限要求的基础路由任何人都可以访问
 */
export const constantRoutes = [{
  path: '/login',
  component: () => import('@/views/login/index'),
  hidden: true,
  meta: {
    title: '后台登录', // 设置该路由在侧边栏和面包屑中展示的名字
  }
},
{
  path: '/404',
  component: () => import('@/views/error-page/404'),
  hidden: true
},
{
  path: '/401',
  component: () => import('@/views/error-page/401'),
  hidden: true
},
// {
//   path: '/',
//   component: Layout,
//   redirect: '/dashboard',
//   children: [{
//     path: 'dashboard',
//     component: () => import('@/views/dashboard/index'),
//     name: 'dashboard',
//     meta: {
//       title: '系统首页', // 设置该路由在侧边栏和面包屑中展示的名字
//       icon: 'el-icon-s-platform', // 设置该路由的图标，支持 svg-class，也支持 el-icon-x element-ui 的 icon
//       affix: true, // 如果设置为true，它则会固定在tags-view中(默认 false)
//       breadcrumb: true //  如果设置为false，则不会在breadcrumb面包屑中显示(默认 true)
//     }
//   }]
// }
]

/**
 * asyncRoutes
  需要权限限制的路由
 */
export const asyncRoutes = [{
  path: '/user',
  component: Layout,
  redirect: '/user/newPwd',
  hidden: true,
  children: [{
    path: 'newPwd',
    component: () => import('@/views/user/newPwd'),
    name: '修改密码',
    meta: {
      title: '修改密码', // 设置该路由在侧边栏和面包屑中展示的名字
      breadcrumb: true //  如果设置为false，则不会在breadcrumb面包屑中显示(默认 true)
    }
  }]
},{
  path: '/',
  component: Layout,
  redirect: '/index',
  alwaysShow: false, // true，这样它就会忽略之前定义的规则，一直显示根路由
  children: [{
    path: 'index',
    component: () => import('@/views/index'),
    name: '系统首页',
    meta: {
      icon: 'el-icon-s-home',
      title: '系统首页',
      roles: ['admin'], // 或者你可以给每一个子路由设置自己的权限
      affix: true,
    }
  }]
},{
  path: '/apply',
  component: Layout,
  redirect: '/apply/index',
  alwaysShow: false, // true，这样它就会忽略之前定义的规则，一直显示根路由
  children: [{
    path: 'index',
    component: () => import('@/views/table/apply'),
    name: '应用管理',
    meta: {
      icon: 'component',
      title: '应用管理',
      roles: ['admin'], // 或者你可以给每一个子路由设置自己的权限
      affix: true,
    }
  }]
},
{
  path: '/users',
  component: Layout,
  redirect: '/users/index',
  children: [{
    path: 'index',
    component: () => import('@/views/table/user'),
    name: '用户管理',
    meta: {
      roles: ['admin'], // 或者你可以给每一个子路由设置自己的权限
      icon: 'user',
      title: '用户管理', // 设置该路由在侧边栏和面包屑中展示的名字
      breadcrumb: true //  如果设置为false，则不会在breadcrumb面包屑中显示(默认 true)
    }
  }]
},
{
  path: '/key',
  component: Layout,
  redirect: '/key/index',
  children: [{
    path: 'index',
    component: () => import('@/views/table/key'),
    name: '卡密管理',
    meta: {
      roles: ['admin'], // 或者你可以给每一个子路由设置自己的权限
      icon: 'password',
      title: '卡密管理', // 设置该路由在侧边栏和面包屑中展示的名字
      breadcrumb: true //  如果设置为false，则不会在breadcrumb面包屑中显示(默认 true)
    }
  }]
}, {
  path: '/power',
  component: Layout,
  redirect: '/power/index',
  children: [{
    path: 'index',
    component: () => import('@/views/table/power'),
    name: '权限管理',
    meta: {
      roles: ['admin'], // 或者你可以给每一个子路由设置自己的权限
      icon: 'el-icon-s-help',
      title: '权限管理', // 设置该路由在侧边栏和面包屑中展示的名字
      breadcrumb: true //  如果设置为false，则不会在breadcrumb面包屑中显示(默认 true)
    }
  }]
},
{
  path: '*',
  redirect: '/404',
  hidden: true
}
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({
    y: 0
  }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
