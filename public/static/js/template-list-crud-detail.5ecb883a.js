"use strict";(self["webpackChunk_6143"]=self["webpackChunk_6143"]||[]).push([[3630],{18824:function(e,t,l){l.r(t),l.d(t,{default:function(){return o}});var u=l(66252);const a=(0,u.Uk)("保存");function i(e,t,l,i,r,d){const n=(0,u.up)("sc-page-header"),o=(0,u.up)("el-alert"),m=(0,u.up)("el-input"),p=(0,u.up)("el-form-item"),s=(0,u.up)("el-button"),c=(0,u.up)("el-form"),f=(0,u.up)("el-card"),w=(0,u.up)("el-main");return(0,u.wg)(),(0,u.iD)(u.HY,null,[(0,u.Wm)(n,{title:r.id?"编辑":"新增",description:"可用于非常复杂的表单提交，如一些较为简单的表单提交应使用dialog或者drawer更合适",icon:"el-icon-burger"},null,8,["title"]),(0,u.Wm)(w,null,{default:(0,u.w5)((()=>[(0,u.Wm)(o,{title:"注意: 因为keep-alive只接受组件name,导致多路由共用组件时,关闭或刷新一个标签导致其他同一组件的页面缓存失效,后续还在寻找完美的解决方案.建议在列表页使用dialog或者drawer形式",type:"error",style:{"margin-bottom":"15px"}}),(0,u.Wm)(f,{shadow:"never"},{default:(0,u.w5)((()=>[(0,u.Wm)(c,{ref:"form","label-width":"100px"},{default:(0,u.w5)((()=>[(0,u.Wm)(p,{label:"id"},{default:(0,u.w5)((()=>[(0,u.Wm)(m,{modelValue:r.id,"onUpdate:modelValue":t[0]||(t[0]=e=>r.id=e)},null,8,["modelValue"])])),_:1}),(0,u.Wm)(p,null,{default:(0,u.w5)((()=>[(0,u.Wm)(s,{type:"primary"},{default:(0,u.w5)((()=>[a])),_:1})])),_:1})])),_:1},512)])),_:1})])),_:1})],64)}var r={name:"listCrud-detail",data(){return{id:this.$route.query.id,input:""}},created(){},mounted(){this.$store.commit("updateViewTagsTitle",this.id?`CURD编辑ID:${this.id}`:"CURD新增")},methods:{}},d=l(83744);const n=(0,d.Z)(r,[["render",i]]);var o=n}}]);