/* empty css             *//* empty css                      *//* empty css                  */import"./el-form-item-4ed993c7.js";/* empty css                 *//* empty css                        *//* empty css                     *//* empty css                    *//* empty css                  */import{d as H,r as o,a as b,o as T,t as z,b as e,w as t,j as r,g as y,h as p,x as J,T as Q,R as X,am as ee,a0 as ae,an as te,M as _,ao as le,m as oe,Y as ne,Z as se,a1 as re,k as de,l as ie,E as ue,a4 as ce,ap as me,aq as pe,ar as _e}from"./index-6dee63ee.js";import fe from"./Menus-3388d1a2.js";import ge from"./Apis-701e4bb5.js";/* empty css                */const ve={class:"authority"},we={class:"table-box"},be={class:"button-left"},ye={class:"pagination"},he={class:"dialog-footer"},Me=H({__name:"index",setup(Ce){const h=o(0),d=o(1),i=o(10),C=o([]),F=o(!1),P=o(!1),B=o(!1),v=o();o();const u=o(!1);o(!1);const f=o(!1),g=o({name:"",age:0}),D=b({name:[{required:!0,trigger:"blur"}]}),S=a=>{i.value=a,c()},U=a=>{d.value=a,c()},c=async()=>{const a=await ee({pageSize:i.value,currentPage:d.value});a.data.code===1e3&&(C.value=a.data.data.data,h.value=a.data.data.total,i.value=a.data.data.pageSize,d.value=a.data.data.currentPage)};c();const M=()=>{},k=(a,l)=>{g.value[a]=l},N=o(null),A=o(null),I=async a=>{ae.confirm("此操作将永久删除该角色, 是否继续?","Warning",{confirmButtonText:"OK",cancelButtonText:"Cancel",type:"warning"}).then(async()=>{(await te({name:a.name})).data.code===1e3?(_({type:"success",message:"角色删除成功"}),await c()):_({type:"error",message:"角色删除失败"})}).catch(()=>{_({type:"info",message:"取消删除角色"})})},j=()=>{u.value=!0},q=async a=>{f.value=!0,g.value=a},s=b({name:"",desc:""});b({name:"",desc:""});const x=async()=>{v.value.resetFields(),u.value=!1},K=async()=>{v.value.validate(async a=>{a&&((await le(s)).data.code===1e3?(_({type:"success",message:"角色创建成功"}),await c(),x()):_({type:"error",message:"角色创建失败"}),u.value=!1)})};return(a,l)=>{const m=oe,w=ne,$=se,L=re,V=de,E=ie,O=ue,W=ce,R=me,Y=pe,Z=_e;return T(),z(X,null,[e(Q,{mode:"out-in",name:"el-fade-in-linear"},{default:t(()=>[r("div",ve,[r("div",we,[r("div",be,[e(m,{type:"primary",icon:"plus",onClick:y(j,["prevent"])},{default:t(()=>[p("新增角色")]),_:1},8,["onClick"])]),e($,{data:C.value,style:{width:"100%"}},{default:t(()=>[e(w,{align:"left",prop:"name",label:"角色名","min-width":"150"}),e(w,{align:"left",prop:"desc",label:"描述","min-width":"180"}),e(w,{align:"left",label:"操作","min-width":"200"},{default:t(n=>[e(m,{icon:"setting",type:"primary",link:"",onClick:y(G=>q(n.row),["prevent"])},{default:t(()=>[p("设置权限")]),_:2},1032,["onClick"]),e(m,{icon:"delete",type:"primary",link:"",onClick:y(G=>I(n.row),["prevent"])},{default:t(()=>[p("删除")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"]),r("div",ye,[e(L,{"current-page":d.value,"onUpdate:currentPage":l[0]||(l[0]=n=>d.value=n),"page-size":i.value,"onUpdate:pageSize":l[1]||(l[1]=n=>i.value=n),"page-sizes":[10,20,50],small:F.value,disabled:B.value,background:P.value,layout:"total, sizes, prev, pager, next, jumper",total:h.value,onSizeChange:S,onCurrentChange:U},null,8,["current-page","page-size","small","disabled","background","total"])]),e(W,{modelValue:u.value,"onUpdate:modelValue":l[4]||(l[4]=n=>u.value=n),title:"创建角色",width:"40%"},{footer:t(()=>[r("span",he,[e(m,{onClick:x},{default:t(()=>[p("关闭")]),_:1}),e(m,{type:"primary",onClick:K},{default:t(()=>[p("提交")]),_:1})])]),default:t(()=>[r("div",null,[e(O,{ref_key:"addFormRef",ref:v,model:s,"status-icon":"",rules:D,"label-width":"120px",style:{"max-width":"380px"},class:"demo-ruleForm"},{default:t(()=>[e(E,{label:"角色名",prop:"name"},{default:t(()=>[e(V,{modelValue:s.name,"onUpdate:modelValue":l[2]||(l[2]=n=>s.name=n),maxlength:"30","show-word-limit":"",placeholder:"请输入角色名"},null,8,["modelValue"])]),_:1}),e(E,{label:"角色描述",prop:"desc"},{default:t(()=>[e(V,{modelValue:s.desc,"onUpdate:modelValue":l[3]||(l[3]=n=>s.desc=n),maxlength:"50",rows:5,type:"textarea","show-word-limit":"",placeholder:"请输入角色描述"},null,8,["modelValue"])]),_:1})]),_:1},8,["model","rules"])])]),_:1},8,["modelValue"]),f.value?(T(),z(Z,{key:0,modelValue:f.value,"onUpdate:modelValue":l[5]||(l[5]=n=>f.value=n),"custom-class":"auth-drawer","with-header":!1,size:"40%",title:"角色配置"},{default:t(()=>[e(Y,{"before-leave":M,type:"border-card"},{default:t(()=>[e(R,{label:"菜单权限"},{default:t(()=>[e(fe,{ref_key:"menus",ref:N,row:g.value,onChangeRow:k},null,8,["row"])]),_:1}),e(R,{label:"API权限"},{default:t(()=>[e(ge,{ref_key:"apis",ref:A,row:g.value,onChangeRow:k},null,8,["row"])]),_:1})]),_:1})]),_:1},8,["modelValue"])):J("",!0)])])]),_:1})],1024)}}});export{Me as default};
