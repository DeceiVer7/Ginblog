import Vue from 'vue'
import VueRouter from 'vue-router'

const ArticleList = () =>
  import(/* webpackChunkName: "group-index" */ '../components/ArticleList.vue')
const Detail = () =>
  import(/* webpackChunkName: "group-detail" */ '../components/Details.vue')
const Category = () =>
  import(/* webpackChunkName: "group-category" */ '../components/CateList.vue')
const Search = () =>
  import(/* webpackChunkName: "group-search" */ '../components/Search.vue')
const AddArt = () =>
  import(/* webpackChunkName: "group-addart" */ '../components/AddArt.vue')
const ArtList = () =>
  import(/* webpackChunkName: "group-artlist" */ '../components/ArtList.vue')
const Profile = () =>
  import(/* webpackChunkName: "group-profile" */ '../components/Profile.vue')
const LikeArtList = () =>
  import(/* webpackChunkName: "group-profile" */ '../components/LikeArtList.vue')
const StarArtList = () =>
  import(/* webpackChunkName: "group-profile" */ '../components/StarArtList.vue')

Vue.use(VueRouter)

//获取原型对象上的push函数
const originalPush = VueRouter.prototype.push
//修改原型对象中的push方法
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

const routes = [
  { path: '/', component: ArticleList, meta: { title: '欢迎来到GinBlog' } },
  {
    path: '/article/detail/:id',
    component: Detail,
    meta: { title: window.sessionStorage.getItem('title') },
    props: true
  },
  {
    path: '/category/:cid',
    component: Category,
    meta: { title: '分类信息' },
    props: true
  },
  {
    path: '/search/:title',
    component: Search,
    meta: { title: '搜索结果' },
    props: true
  },
  {
    path: '/addart',
    component: AddArt,
    meta: { title: '写文章' },
    props: true
  },
  {
    path: '/addart/:id',
    component: AddArt,
    meta: { title: '编辑文章' },
    props: true
  },
  {
    path: '/artlist/:author',
    name: 'Artlist',
    component: ArtList,
    meta: { title: '文章管理' },
    props: true
  },
  {
    path: '/profile',
    component: Profile,
    meta: { title: '个人信息' },
    props: true
  },
  {
    path: '/likeartlist',
    component: LikeArtList,
    meta: { title: '点赞文章列表' },
    props: true
  },
  {
    path: '/starartlist',
    component: StarArtList,
    meta: { title: '点赞文章列表' },
    props: true
  },
]

const router = new VueRouter({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title ? to.meta.title : '加载中'
  }
  next()
})

export default router
