/* empty css                  *//* empty css                        *//* empty css                 *//* empty css                   */import"./el-form-item-4ed993c7.js";/* empty css                    */import{X as b,d as K,r as i,a as C,o as R,c as D,j as g,b as a,w as t,h as m,g as E,F as W,s as ee,L as h,m as ae,Y as le,Z as te,$ as oe,k as se,l as ne,a0 as re,a1 as de,E as ue,a2 as pe}from"./index-95ad95c9.js";const ie=u=>b({url:"/instance/list",method:"post",data:u}),me=u=>b({url:"/instance",method:"post",data:u}),ce=u=>b({url:"/instance/add",method:"post",data:u}),fe=u=>b({url:"/instance",method:"patch",data:u}),ge=u=>b({url:"/instance",method:"delete",data:u}),be=u=>b({url:"/instance/ping",method:"post",data:u}),ve={class:"table-box"},_e={class:"button-left"},we={class:"pagination"},ye={class:"dialog-footer"},Ve={class:"dialog-footer"},qe=K({__name:"index",setup(u){const P=i(0),v=i(1),_=i(10),F=i([]),B=i(!1),L=i(!1),M=i(!1);i("");const T=[{value:"mysql",label:"MySQL"},{value:"redis",label:"Redis"},{value:"mongodb",label:"MongoDB"},{value:"elasticsearch",label:"Elasticsearch"},{value:"postgresql",label:"PostgreSQL"}],c=async()=>{const o=await ie({pageSize:_.value,currentPage:v.value});o.data.code===1e3&&(F.value=o.data.data.data,P.value=o.data.data.total,_.value=o.data.data.pageSize,v.value=o.data.data.currentPage)};(async()=>{await c()})();const $=o=>{_.value=o,c()},N=o=>{v.value=o,c()},w=i(!1),y=i(!1),j=async o=>{(await ge({name:o.name})).status===200&&(h.success("删除成功"),await c())},Q=()=>{w.value=!0},O=async o=>{y.value=!0;const e=await me({name:o.name});e.status===200&&(s.name=e.data.data[0].name,s.db_type=e.data.data[0].db_type,s.host=e.data.data[0].host,s.port=e.data.data[0].port,s.user=e.data.data[0].user,s.password=e.data.data[0].password)},X=async o=>{(await be({name:o.name})).status===200?h({type:"success",message:"连接数据库成功"}):h({type:"error",message:"连接数据库失败"})},U=C({name:[{required:!0,trigger:"blur"}],db_type:[{required:!0,trigger:"blur"}],host:[{required:!0,trigger:"blur"}],port:[{required:!0,trigger:"blur"}],user:[{required:!0,trigger:"blur"}],password:[{validator:(o,e,r)=>{if(e==="")r(new Error("请输入密码"));else{if(n.checkPass!==""){if(!V.value)return;V.value.validateField("checkPass",()=>null)}r()}}},{required:!0,min:6,trigger:"blur"}],checkPass:[{validator:(o,e,r)=>{e===""?r(new Error("请再次输入密码")):e!=n.password?r(new Error("输入的密码不相同!")):r()}},{required:!0,min:6,trigger:"blur"}]}),V=i(),k=i(),s=C({name:"",db_type:"",host:"",port:"",user:"",password:""}),n=C({name:"",db_type:"",host:"",port:"",user:"",password:"",checkPass:""}),x=async()=>{V.value.resetFields(),w.value=!1},Y=async()=>{V.value.validate(async o=>{o&&((await ce(n)).status===200&&(h({type:"success",message:"实例创建成功"}),await c(),x()),w.value=!1)})},z=async()=>{k.value.resetFields(),y.value=!1},Z=async()=>{k.value.validate(async o=>{o&&((await fe(s)).status===200&&(h({type:"success",message:"实例修改成功"}),await c(),z()),y.value=!1)})};return(o,e)=>{const r=ae,f=le,A=te,G=oe,p=se,d=ne,H=re,J=de,I=ue,q=pe;return R(),D("div",ve,[g("div",_e,[a(r,{type:"primary",icon:"plus",onClick:Q},{default:t(()=>[m("新增实例")]),_:1})]),a(A,{data:F.value,style:{width:"100%"}},{default:t(()=>[a(f,{align:"left",prop:"name",label:"实例名"}),a(f,{align:"left",prop:"db_type",label:"数据库类型"}),a(f,{align:"left",prop:"host",label:"数据库地址"}),a(f,{align:"left",prop:"port",label:"数据库端口"}),a(f,{align:"left",prop:"user",label:"连接用户"}),a(f,{align:"left",label:"操作",width:"180"},{default:t(l=>[a(r,{link:"",type:"primary",size:"small",onClick:E(S=>O(l.row),["prevent"])},{default:t(()=>[m(" 编辑 ")]),_:2},1032,["onClick"]),a(r,{link:"",type:"primary",size:"small",onClick:E(S=>j(l.row),["prevent"])},{default:t(()=>[m(" 删除 ")]),_:2},1032,["onClick"]),a(r,{link:"",type:"primary",size:"small",onClick:E(S=>X(l.row),["prevent"])},{default:t(()=>[m(" 连接测试 ")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"]),g("div",we,[a(G,{"current-page":v.value,"onUpdate:currentPage":e[0]||(e[0]=l=>v.value=l),"page-size":_.value,"onUpdate:pageSize":e[1]||(e[1]=l=>_.value=l),"page-sizes":[10,20,50],small:B.value,disabled:M.value,background:L.value,layout:"total, sizes, prev, pager, next, jumper",total:P.value,onSizeChange:$,onCurrentChange:N},null,8,["current-page","page-size","small","disabled","background","total"])]),a(q,{modelValue:w.value,"onUpdate:modelValue":e[9]||(e[9]=l=>w.value=l),title:"创建实例",width:"40%"},{footer:t(()=>[g("span",ye,[a(r,{onClick:x},{default:t(()=>[m("关闭")]),_:1}),a(r,{type:"primary",onClick:Y},{default:t(()=>[m("提交")]),_:1})])]),default:t(()=>[g("div",null,[a(I,{ref_key:"insFormRef",ref:V,model:n,"status-icon":"",rules:U,"label-width":"120px",style:{"max-width":"380px"},class:"demo-ruleForm"},{default:t(()=>[a(d,{label:"实例名",prop:"name"},{default:t(()=>[a(p,{modelValue:n.name,"onUpdate:modelValue":e[2]||(e[2]=l=>n.name=l),maxlength:"30","show-word-limit":"",placeholder:"请输入实例名"},null,8,["modelValue"])]),_:1}),a(d,{label:"数据库类型",prop:"db_type"},{default:t(()=>[a(J,{modelValue:n.db_type,"onUpdate:modelValue":e[3]||(e[3]=l=>n.db_type=l),placeholder:"数据库类型"},{default:t(()=>[(R(),D(W,null,ee(T,l=>a(H,{key:l.value,label:l.label,value:l.value},null,8,["label","value"])),64))]),_:1},8,["modelValue"])]),_:1}),a(d,{label:"数据库地址",prop:"host"},{default:t(()=>[a(p,{modelValue:n.host,"onUpdate:modelValue":e[4]||(e[4]=l=>n.host=l),placeholder:"请输入数据库地址"},null,8,["modelValue"])]),_:1}),a(d,{label:"数据库端口",prop:"port"},{default:t(()=>[a(p,{modelValue:n.port,"onUpdate:modelValue":e[5]||(e[5]=l=>n.port=l),placeholder:"请输入数据库端口"},null,8,["modelValue"])]),_:1}),a(d,{label:"数据库用户",prop:"user"},{default:t(()=>[a(p,{modelValue:n.user,"onUpdate:modelValue":e[6]||(e[6]=l=>n.user=l),placeholder:"请输入连接用户"},null,8,["modelValue"])]),_:1}),a(d,{label:"密码",prop:"password"},{default:t(()=>[a(p,{modelValue:n.password,"onUpdate:modelValue":e[7]||(e[7]=l=>n.password=l),placeholder:"请输入密码",autocomplete:"off",type:"password","show-password":""},null,8,["modelValue"])]),_:1}),a(d,{label:"确认密码",prop:"checkPass"},{default:t(()=>[a(p,{modelValue:n.checkPass,"onUpdate:modelValue":e[8]||(e[8]=l=>n.checkPass=l),placeholder:"请再次输入密码",autocomplete:"off",type:"password","show-password":""},null,8,["modelValue"])]),_:1})]),_:1},8,["model","rules"])])]),_:1},8,["modelValue"]),a(q,{modelValue:y.value,"onUpdate:modelValue":e[16]||(e[16]=l=>y.value=l),title:"修改实例",width:"40%"},{footer:t(()=>[g("span",Ve,[a(r,{onClick:z},{default:t(()=>[m("关闭")]),_:1}),a(r,{type:"primary",onClick:Z},{default:t(()=>[m("提交")]),_:1})])]),default:t(()=>[g("div",null,[a(I,{ref_key:"editFormRef",ref:k,model:s,"status-icon":"",rules:U,"label-width":"120px",style:{"max-width":"380px"},class:"demo-ruleForm"},{default:t(()=>[a(d,{label:"实例名",prop:"name"},{default:t(()=>[a(p,{modelValue:s.name,"onUpdate:modelValue":e[10]||(e[10]=l=>s.name=l),disabled:"",maxlength:"50","show-word-limit":"",placeholder:"请输入实例名"},null,8,["modelValue"])]),_:1}),a(d,{label:"数据库类型",prop:"db_type"},{default:t(()=>[a(p,{modelValue:s.db_type,"onUpdate:modelValue":e[11]||(e[11]=l=>s.db_type=l),disabled:"",placeholder:"请输入数据库类型"},null,8,["modelValue"])]),_:1}),a(d,{label:"数据库地址",prop:"host"},{default:t(()=>[a(p,{modelValue:s.host,"onUpdate:modelValue":e[12]||(e[12]=l=>s.host=l),placeholder:"请输入数据库地址"},null,8,["modelValue"])]),_:1}),a(d,{label:"数据库端口",prop:"port"},{default:t(()=>[a(p,{modelValue:s.port,"onUpdate:modelValue":e[13]||(e[13]=l=>s.port=l),placeholder:"请输入数据库端口"},null,8,["modelValue"])]),_:1}),a(d,{label:"数据库用户",prop:"user"},{default:t(()=>[a(p,{modelValue:s.user,"onUpdate:modelValue":e[14]||(e[14]=l=>s.user=l),placeholder:"请输入连接用户"},null,8,["modelValue"])]),_:1}),a(d,{label:"密码",prop:"password"},{default:t(()=>[a(p,{modelValue:s.password,"onUpdate:modelValue":e[15]||(e[15]=l=>s.password=l),placeholder:"请输入密码",autocomplete:"off",type:"password","show-password":""},null,8,["modelValue"])]),_:1})]),_:1},8,["model","rules"])])]),_:1},8,["modelValue"])])}}});export{qe as default};
