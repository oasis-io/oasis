/* empty css             *//* empty css                  */import"./el-form-item-4ed993c7.js";/* empty css                 */import{d as x,r as w,a as u,u as y,o as b,c as h,b as e,w as s,e as v,f as V,E as k,g as I,h as E,p as F,i as S,j as c,k as B,l as C,m as K,_ as N}from"./index-6dee63ee.js";const m=t=>(F("data-v-9cc40884"),t=t(),S(),t),U={class:"login-container"},q=m(()=>c("div",{class:"login-background"},null,-1)),R=m(()=>c("div",{class:"login-title"},"Oasis 数据库运维平台",-1)),j=x({__name:"index",setup(t){const n=w(),o=u({username:"",password:""}),_=u({username:[{required:!0,message:"请输入用户名",trigger:"blur"}],password:[{required:!0,message:"请输入密码",trigger:"blur"}]}),g=y(),i=()=>{n.value.validate(async p=>{if(p)await g.LoginIn(o);else return!1})};return(p,a)=>{const d=B,l=C,f=K;return b(),h("div",U,[q,R,e(V(k),{model:o,rules:_,ref_key:"loginFormRef",ref:n,"label-width":"auto",class:"login-form",onKeyup:v(i,["enter"])},{default:s(()=>[e(l,{label:"",prop:"username",style:{"margin-bottom":"20px"}},{default:s(()=>[e(d,{modelValue:o.username,"onUpdate:modelValue":a[0]||(a[0]=r=>o.username=r),placeholder:"用户名",class:"login-input",maxlength:"30","prefix-icon":"el-icon-user",style:{"--el-input-border-radius":"25px"}},null,8,["modelValue"])]),_:1}),e(l,{label:"",prop:"password",style:{"margin-top":"0"}},{default:s(()=>[e(d,{type:"password",modelValue:o.password,"onUpdate:modelValue":a[1]||(a[1]=r=>o.password=r),placeholder:"密码",class:"login-input",maxlength:"30","prefix-icon":"el-icon-lock",style:{"--el-input-border-radius":"25px"},"show-password":""},null,8,["modelValue"])]),_:1}),e(l,{style:{"margin-top":"20px"}},{default:s(()=>[e(f,{type:"primary",onClick:I(i,["prevent"]),class:"login-button"},{default:s(()=>[E("登录")]),_:1},8,["onClick"])]),_:1})]),_:1},8,["model","rules","onKeyup"])])}}});const A=N(j,[["__scopeId","data-v-9cc40884"]]);export{A as default};