"use strict";(self["webpackChunk_6143"]=self["webpackChunk_6143"]||[]).push([[1686],{87343:function(t,e,n){n.r(e),n.d(e,{default:function(){return g}});var i=n(66252);const o=t=>((0,i.dD)("data-v-03a2c1a0"),t=t(),(0,i.Cn)(),t),r=(0,i.Uk)("普通打印"),l=o((()=>(0,i._)("div",{style:{height:"20px"}},null,-1))),c={class:"printMain",ref:"printMain"},a={class:"item"},s=(0,i.Uk)("打印内容1 "),u=(0,i.Uk)(),d=o((()=>(0,i._)("p",{class:"no-print"},"忽略打印",-1))),p=o((()=>(0,i._)("div",{style:{"page-break-after":"always"}},null,-1))),m=o((()=>(0,i._)("div",{class:"item"},"打印内容2",-1))),f=(0,i.Uk)("动态打印");function h(t,e,n,o,h,y){const v=(0,i.up)("el-alert"),b=(0,i.up)("el-button"),k=(0,i.up)("el-icon-eleme-filled"),w=(0,i.up)("el-icon"),g=(0,i.up)("el-tab-pane"),_=(0,i.up)("el-tabs"),M=(0,i.up)("el-card"),T=(0,i.up)("el-main");return(0,i.wg)(),(0,i.j4)(T,null,{default:(0,i.w5)((()=>[(0,i.Wm)(M,{shadow:"never"},{default:(0,i.w5)((()=>[(0,i.Wm)(_,{"tab-position":"top"},{default:(0,i.w5)((()=>[(0,i.Wm)(g,{label:"普通打印"},{default:(0,i.w5)((()=>[(0,i.Wm)(v,{title:"打印当前页面已存在的元素,如包含.no-print样式就忽略,分页打印就需要{page-break-after: always}控制",type:"success",style:{"margin-bottom":"20px"}}),(0,i.Wm)(b,{type:"primary",onClick:y.print},{default:(0,i.w5)((()=>[r])),_:1},8,["onClick"]),l,(0,i._)("div",c,[(0,i._)("div",a,[s,(0,i.Wm)(w,null,{default:(0,i.w5)((()=>[(0,i.Wm)(k)])),_:1}),u,d]),p,m],512)])),_:1}),(0,i.Wm)(g,{label:"动态打印"},{default:(0,i.w5)((()=>[(0,i.Wm)(v,{title:"打印创建的DOM结构,适用于远程获取模板后打印",type:"success",style:{"margin-bottom":"20px"}}),(0,i.Wm)(b,{type:"primary",onClick:y.print2},{default:(0,i.w5)((()=>[f])),_:1},8,["onClick"])])),_:1})])),_:1})])),_:1})])),_:1})}const y=function(t,e){if(!(this instanceof y))return new y(t,e);if(this.options=this.extend({noPrint:".no-print"},e),"string"===typeof t)try{this.dom=document.querySelector(t)}catch{var n=document.createElement("div");n.innerHTML=t,this.dom=n}else this.isDOM(t),this.dom=this.isDOM(t)?t:t.$el;this.init()};y.prototype={init:function(){var t=this.getStyle()+this.getHtml();this.writeIframe(t)},extend:function(t,e){for(var n in e)t[n]=e[n];return t},getStyle:function(){for(var t="",e=document.querySelectorAll("style,link"),n=0;n<e.length;n++)t+=e[n].outerHTML;return t+="<style>"+(this.options.noPrint?this.options.noPrint:".no-print")+"{display:none;}</style>",t+="<style>html,body{background-color:#fff;}</style>",t},getHtml:function(){for(var t=document.querySelectorAll("input"),e=document.querySelectorAll("textarea"),n=document.querySelectorAll("select"),i=0;i<t.length;i++)"checkbox"==t[i].type||"radio"==t[i].type?1==t[i].checked?t[i].setAttribute("checked","checked"):t[i].removeAttribute("checked"):(t[i].type,t[i].setAttribute("value",t[i].value));for(var o=0;o<e.length;o++)"textarea"==e[o].type&&(e[o].innerHTML=e[o].value);for(var r=0;r<n.length;r++)if("select-one"==n[r].type){var l=n[r].children;for(var c in l)"OPTION"==l[c].tagName&&(1==l[c].selected?l[c].setAttribute("selected","selected"):l[c].removeAttribute("selected"))}return this.dom.outerHTML},writeIframe:function(t){var e,n,i=document.createElement("iframe"),o=document.body.appendChild(i);i.id="myIframe",i.setAttribute("style","position:absolute;width:0;height:0;top:-10px;left:-10px;"),e=o.contentWindow||o.contentDocument,n=o.contentDocument||o.contentWindow.document,n.open(),n.write(t),n.close();var r=this;i.onload=function(){r.toPrint(e),setTimeout((function(){document.body.removeChild(i)}),100)}},toPrint:function(t){try{setTimeout((function(){t.focus();try{t.document.execCommand("print",!1,null)||t.print()}catch(e){t.print()}t.close()}),10)}catch(e){console.log("err",e)}},isDOM:"object"===typeof HTMLElement?function(t){return t instanceof HTMLElement}:function(t){return t&&"object"===typeof t&&1===t.nodeType&&"string"===typeof t.nodeName}};var v=y,b={name:"print",data(){return{}},methods:{print(){v(this.$refs.printMain)},print2(){var t="<h2>TITLE</h2><p>后创建的DOM结构</p>";v(t)}}},k=n(83744);const w=(0,k.Z)(b,[["render",h],["__scopeId","data-v-03a2c1a0"]]);var g=w}}]);