"use strict";(self["webpackChunk_6143"]=self["webpackChunk_6143"]||[]).push([[976],{18720:function(e,l,t){t.r(l),t.d(l,{default:function(){return b}});var a=t(66252);const i={class:"left-panel"},s=(0,a.Uk)("分配角色"),o=(0,a.Uk)("密码重置"),n={class:"right-panel"},r={class:"right-panel-search"},d=(0,a.Uk)("查看"),u=(0,a.Uk)("编辑"),c=(0,a.Uk)("删除");function p(e,l,t,p,h,m){const g=(0,a.up)("el-input"),f=(0,a.up)("el-header"),b=(0,a.up)("el-tree"),w=(0,a.up)("el-main"),k=(0,a.up)("el-container"),_=(0,a.up)("el-aside"),W=(0,a.up)("el-button"),v=(0,a.up)("el-table-column"),y=(0,a.up)("el-avatar"),$=(0,a.up)("el-popconfirm"),C=(0,a.up)("el-button-group"),x=(0,a.up)("scTable"),D=(0,a.up)("save-dialog"),T=(0,a.Q2)("loading");return(0,a.wg)(),(0,a.iD)(a.HY,null,[(0,a.Wm)(k,null,{default:(0,a.w5)((()=>[(0,a.wy)(((0,a.wg)(),(0,a.j4)(_,{width:"200px"},{default:(0,a.w5)((()=>[(0,a.Wm)(k,null,{default:(0,a.w5)((()=>[(0,a.Wm)(f,null,{default:(0,a.w5)((()=>[(0,a.Wm)(g,{placeholder:"输入关键字进行过滤",modelValue:h.groupFilterText,"onUpdate:modelValue":l[0]||(l[0]=e=>h.groupFilterText=e),clearable:""},null,8,["modelValue"])])),_:1}),(0,a.Wm)(w,{class:"nopadding"},{default:(0,a.w5)((()=>[(0,a.Wm)(b,{ref:"group",class:"menu","node-key":"id",data:h.group,"current-node-key":"","highlight-current":!0,"expand-on-click-node":!1,"filter-node-method":m.groupFilterNode,onNodeClick:m.groupClick},null,8,["data","filter-node-method","onNodeClick"])])),_:1})])),_:1})])),_:1})),[[T,h.showGrouploading]]),(0,a.Wm)(k,null,{default:(0,a.w5)((()=>[(0,a.Wm)(f,null,{default:(0,a.w5)((()=>[(0,a._)("div",i,[(0,a.Wm)(W,{type:"primary",icon:"el-icon-plus",onClick:m.add},null,8,["onClick"]),(0,a.Wm)(W,{type:"danger",plain:"",icon:"el-icon-delete",disabled:0==h.selection.length,onClick:m.batch_del},null,8,["disabled","onClick"]),(0,a.Wm)(W,{type:"primary",plain:"",disabled:0==h.selection.length},{default:(0,a.w5)((()=>[s])),_:1},8,["disabled"]),(0,a.Wm)(W,{type:"primary",plain:"",disabled:0==h.selection.length},{default:(0,a.w5)((()=>[o])),_:1},8,["disabled"])]),(0,a._)("div",n,[(0,a._)("div",r,[(0,a.Wm)(g,{modelValue:h.search.name,"onUpdate:modelValue":l[1]||(l[1]=e=>h.search.name=e),placeholder:"登录账号 / 姓名",clearable:""},null,8,["modelValue"]),(0,a.Wm)(W,{type:"primary",icon:"el-icon-search",onClick:m.upsearch},null,8,["onClick"])])])])),_:1}),(0,a.Wm)(w,{class:"nopadding"},{default:(0,a.w5)((()=>[(0,a.Wm)(x,{ref:"table",apiObj:h.apiObj,onSelectionChange:m.selectionChange,stripe:"",remoteSort:"",remoteFilter:""},{default:(0,a.w5)((()=>[(0,a.Wm)(v,{type:"selection",width:"50"}),(0,a.Wm)(v,{label:"ID",prop:"id",width:"80",sortable:"custom"}),(0,a.Wm)(v,{label:"头像",width:"80","column-key":"filterAvatar",filters:[{text:"已上传",value:"1"},{text:"未上传",value:"0"}]},{default:(0,a.w5)((e=>[(0,a.Wm)(y,{src:e.row.avatar,size:"small"},null,8,["src"])])),_:1}),(0,a.Wm)(v,{label:"登录账号",prop:"userName",width:"150",sortable:"custom","column-key":"filterUserName",filters:[{text:"系统账号",value:"1"},{text:"普通账号",value:"0"}]}),(0,a.Wm)(v,{label:"姓名",prop:"name",width:"150",sortable:"custom"}),(0,a.Wm)(v,{label:"所属角色",prop:"groupName",width:"200",sortable:"custom"}),(0,a.Wm)(v,{label:"加入时间",prop:"date",width:"170",sortable:"custom"}),(0,a.Wm)(v,{label:"操作",fixed:"right",align:"right",width:"160"},{default:(0,a.w5)((e=>[(0,a.Wm)(C,null,{default:(0,a.w5)((()=>[(0,a.Wm)(W,{text:"",type:"primary",size:"small",onClick:l=>m.table_show(e.row,e.$index)},{default:(0,a.w5)((()=>[d])),_:2},1032,["onClick"]),(0,a.Wm)(W,{text:"",type:"primary",size:"small",onClick:l=>m.table_edit(e.row,e.$index)},{default:(0,a.w5)((()=>[u])),_:2},1032,["onClick"]),(0,a.Wm)($,{title:"确定删除吗？",onConfirm:l=>m.table_del(e.row,e.$index)},{reference:(0,a.w5)((()=>[(0,a.Wm)(W,{text:"",type:"primary",size:"small"},{default:(0,a.w5)((()=>[c])),_:1})])),_:2},1032,["onConfirm"])])),_:2},1024)])),_:1})])),_:1},8,["apiObj","onSelectionChange"])])),_:1})])),_:1})])),_:1}),h.dialog.save?((0,a.wg)(),(0,a.j4)(D,{key:0,ref:"saveDialog",onSuccess:m.handleSuccess,onClosed:l[2]||(l[2]=e=>h.dialog.save=!1)},null,8,["onSuccess"])):(0,a.kq)("",!0)],64)}var h=t(65796),m={name:"user",components:{saveDialog:h["default"]},data(){return{dialog:{save:!1},showGrouploading:!1,groupFilterText:"",group:[],apiObj:this.$API.system.user.list,selection:[],search:{name:null}}},watch:{groupFilterText(e){this.$refs.group.filter(e)}},mounted(){this.getGroup()},methods:{add(){this.dialog.save=!0,this.$nextTick((()=>{this.$refs.saveDialog.open()}))},table_edit(e){this.dialog.save=!0,this.$nextTick((()=>{this.$refs.saveDialog.open("edit").setData(e)}))},table_show(e){this.dialog.save=!0,this.$nextTick((()=>{this.$refs.saveDialog.open("show").setData(e)}))},async table_del(e,l){var t={id:e.id},a=await this.$API.demo.post.post(t);200==a.code?(this.$refs.table.tableData.splice(l,1),this.$message.success("删除成功")):this.$alert(a.message,"提示",{type:"error"})},async batch_del(){this.$confirm(`确定删除选中的 ${this.selection.length} 项吗？`,"提示",{type:"warning"}).then((()=>{const e=this.$loading();this.selection.forEach((e=>{this.$refs.table.tableData.forEach(((l,t)=>{e.id===l.id&&this.$refs.table.tableData.splice(t,1)}))})),e.close(),this.$message.success("操作成功")})).catch((()=>{}))},selectionChange(e){this.selection=e},async getGroup(){this.showGrouploading=!0;var e=await this.$API.system.dept.list.get();this.showGrouploading=!1;var l={id:"",label:"所有"};e.data.unshift(l),this.group=e.data},groupFilterNode(e,l){return!e||-1!==l.label.indexOf(e)},groupClick(e){var l={groupId:e.id};this.$refs.table.reload(l)},upsearch(){this.$refs.table.upData(this.search)},handleSuccess(e,l){"add"==l?(e.id=(new Date).getTime(),this.$refs.table.tableData.unshift(e)):"edit"==l&&this.$refs.table.tableData.filter((l=>l.id===e.id)).forEach((l=>{Object.assign(l,e)}))}}},g=t(83744);const f=(0,g.Z)(m,[["render",p]]);var b=f}}]);