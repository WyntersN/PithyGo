"use strict";(self["webpackChunk_6143"]=self["webpackChunk_6143"]||[]).push([[2577],{42138:function(a,l,e){e.r(l),e.d(l,{default:function(){return W}});var o=e(66252);const i=(0,o.Uk)("默认"),d=(0,o.Uk)("加载"),n=(0,o.Uk)("禁止拖拽最大化和关闭"),t=(0,o.Uk)("异步加载1"),g=(0,o.Uk)("异步加载2"),u=(0,o.Uk)(" 内容 "),m=(0,o.Uk)("取 消"),r=(0,o.Uk)("确 定"),s=(0,o.Uk)("取 消"),p=(0,o.Uk)("确 定"),k=(0,o.Uk)(" 内容 "),y=(0,o.Uk)("取 消"),c=(0,o.Uk)("确 定");function w(a,l,e,w,f,C){const h=(0,o.up)("el-alert"),W=(0,o.up)("el-button"),_=(0,o.up)("el-card"),U=(0,o.up)("el-main"),D=(0,o.up)("sc-dialog"),b=(0,o.up)("el-empty"),V=(0,o.up)("dialog1"),v=(0,o.up)("dialog2");return(0,o.wg)(),(0,o.iD)(o.HY,null,[(0,o.Wm)(U,null,{default:(0,o.w5)((()=>[(0,o.Wm)(h,{title:"二次封装el-dialog,加入加载中/最大化.",type:"success",style:{margin:"0 0 20px 0"}}),(0,o.Wm)(_,{shadow:"never",header:"内置"},{default:(0,o.w5)((()=>[(0,o.Wm)(W,{type:"primary",onClick:C.open1},{default:(0,o.w5)((()=>[i])),_:1},8,["onClick"]),(0,o.Wm)(W,{type:"primary",onClick:C.open2},{default:(0,o.w5)((()=>[d])),_:1},8,["onClick"]),(0,o.Wm)(W,{type:"primary",onClick:C.open3},{default:(0,o.w5)((()=>[n])),_:1},8,["onClick"])])),_:1}),(0,o.Wm)(_,{shadow:"never",header:"异步",style:{"margin-top":"15px"}},{default:(0,o.w5)((()=>[(0,o.Wm)(W,{type:"primary",onClick:C.asyn1},{default:(0,o.w5)((()=>[t])),_:1},8,["onClick"]),(0,o.Wm)(W,{type:"primary",onClick:C.asyn2},{default:(0,o.w5)((()=>[g])),_:1},8,["onClick"]),(0,o.Wm)(h,{title:"适用于页面有很多弹窗操作,利用异步组件按需加载,加快首屏的加载速度和打包体积",style:{"margin-top":"20px"}})])),_:1})])),_:1}),(0,o.Wm)(D,{modelValue:f.dialog1,"onUpdate:modelValue":l[2]||(l[2]=a=>f.dialog1=a),draggable:"",title:"提示"},{footer:(0,o.w5)((()=>[(0,o.Wm)(W,{onClick:l[0]||(l[0]=a=>f.dialog1=!1)},{default:(0,o.w5)((()=>[m])),_:1}),(0,o.Wm)(W,{type:"primary",onClick:l[1]||(l[1]=a=>f.dialog1=!1)},{default:(0,o.w5)((()=>[r])),_:1})])),default:(0,o.w5)((()=>[u])),_:1},8,["modelValue"]),(0,o.Wm)(D,{modelValue:f.dialog2,"onUpdate:modelValue":l[5]||(l[5]=a=>f.dialog2=a),draggable:"",title:"模拟加载",width:400,loading:f.dialog2Loading},{footer:(0,o.w5)((()=>[(0,o.Wm)(W,{onClick:l[3]||(l[3]=a=>f.dialog2=!1)},{default:(0,o.w5)((()=>[s])),_:1}),(0,o.Wm)(W,{type:"primary",onClick:l[4]||(l[4]=a=>f.dialog2=!1),loading:f.dialog2Loading},{default:(0,o.w5)((()=>[p])),_:1},8,["loading"])])),default:(0,o.w5)((()=>[(0,o.Wm)(b,{description:"NO Data","image-size":80})])),_:1},8,["modelValue","loading"]),(0,o.Wm)(D,{modelValue:f.dialog3,"onUpdate:modelValue":l[8]||(l[8]=a=>f.dialog3=a),title:"禁用拖拽","show-fullscreen":!1,"show-close":!1},{footer:(0,o.w5)((()=>[(0,o.Wm)(W,{onClick:l[6]||(l[6]=a=>f.dialog3=!1)},{default:(0,o.w5)((()=>[y])),_:1}),(0,o.Wm)(W,{type:"primary",onClick:l[7]||(l[7]=a=>f.dialog3=!1)},{default:(0,o.w5)((()=>[c])),_:1})])),default:(0,o.w5)((()=>[k])),_:1},8,["modelValue"]),f.asynDialog1?((0,o.wg)(),(0,o.j4)(V,{key:0,draggable:"",onClosed:l[9]||(l[9]=a=>f.asynDialog1=!1)})):(0,o.kq)("",!0),f.asynDialog2?((0,o.wg)(),(0,o.j4)(v,{key:1,draggable:"",onClosed:l[10]||(l[10]=a=>f.asynDialog2=!1)})):(0,o.kq)("",!0)],64)}var f={name:"dialogExtend",components:{dialog1:(0,o.RC)((()=>e.e(3270).then(e.bind(e,4451)))),dialog2:(0,o.RC)((()=>e.e(7666).then(e.bind(e,76690))))},data(){return{dialog1:!1,dialog2:!1,dialog3:!1,dialog2Loading:!1,asynDialog1:!1,asynDialog2:!1}},mounted(){},methods:{open1(){this.dialog1=!0},open2(){this.dialog2=!0,this.dialog2Loading=!0,setTimeout((()=>{this.dialog2Loading=!1}),1e3)},open3(){this.dialog3=!0},asyn1(){this.asynDialog1=!0},asyn2(){this.asynDialog2=!0}}},C=e(83744);const h=(0,C.Z)(f,[["render",w]]);var W=h}}]);