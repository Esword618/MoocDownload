import{d as F,r as p,a as B,o as V,c as k,b as e,w as l,F as y,e as b,f as t,g as C,u as f,s as D,h as H,E as _}from"./index.e54c3ed6.js";const z=b("h1",{cless:"title",style:{color:"black"}},"\u4E3B\u9875",-1),N=H("\u67E5\u8BE2"),P=F({__name:"Home",setup(S){const u=p(""),c=p("");p(!1);const a=B([{}]),d={children:"children",label:"label"};let m=600,v="600px";const w=[{id:1,label:"\u7B2C\u4E00\u7AE0",children:[{id:4,label:"Level two 1-1",children:[{id:9,label:"Level three 1-1-1"},{id:10,label:"Level three 1-1-2"}]}]},{id:2,label:"Level one 2",children:[{id:5,label:"Level two 2-1"},{id:6,label:"Level two 2-2"}]},{id:3,label:"Level one 3",children:[{id:7,label:"Level two 3-1"},{id:8,label:"Level two 3-2"}]},{id:4,label:"Level one 3",children:[{id:9,label:"Level two 3-1"},{id:10,label:"Level two 3-2"}]}],E=()=>{if(u.value==""){_({message:"\u94FE\u63A5\u4E0D\u80FD\u4E3A\u7A7A!",type:"warning"});return}if(c.value==""){_({message:"\u5E73\u53F0\u4E0D\u80FD\u4E3A\u7A7A!",type:"warning"});return}_({message:"\u67E5\u8BE2\u6210\u529F",type:"success"}),Object.assign(a,w),console.log(a)};return V(()=>{window.addEventListener("resize",()=>{m=document.documentElement.clientHeight*600/768,v=m.toString()+"px"})}),(U,o)=>{const n=t("el-col"),h=t("el-button"),r=t("el-option"),g=t("el-select"),x=t("el-input"),i=t("el-tree"),L=t("el-scrollbar"),A=t("el-row");return C(),k(y,null,[z,e(A,null,{default:l(()=>[e(n,{span:1}),e(n,{span:10},{default:l(()=>[b("div",null,[e(x,{clearable:"",size:"default",modelValue:u.value,"onUpdate:modelValue":o[1]||(o[1]=s=>u.value=s),placeholder:"\u8BF7\u628A\u94FE\u63A5\u7C98\u8D34\u6B64\u5904",class:"input-with-select"},{prepend:l(()=>[e(h,{icon:f(D)},null,8,["icon"])]),append:l(()=>[e(g,{modelValue:c.value,"onUpdate:modelValue":o[0]||(o[0]=s=>c.value=s),placeholder:"Select",style:{width:"115px"}},{default:l(()=>[e(r,{label:"mooc",value:"1"}),e(r,{label:"xtzx",value:"2"}),e(r,{label:"bilibili",value:"3"})]),_:1},8,["modelValue"])]),_:1},8,["modelValue"])]),b("div",null,[e(h,{size:"large",onClick:o[2]||(o[2]=s=>E()),round:""},{default:l(()=>[N]),_:1})])]),_:1}),e(n,{span:1}),e(n,{span:10,class:"tree-box"},{default:l(()=>[e(L,{height:f(v)},{default:l(()=>[e(i,{data:a,props:d,"show-checkbox":""},null,8,["data"]),e(i,{class:"tree-box",data:a,props:d,"show-checkbox":""},null,8,["data"]),e(i,{class:"tree-box",data:a,props:d,"default-expand-all":"true","show-checkbox":""},null,8,["data"])]),_:1},8,["height"])]),_:1})]),_:1})],64)}}});export{P as default};