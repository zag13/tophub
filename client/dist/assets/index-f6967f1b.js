import{u as i,_ as p,w as C,a as v,E as S,b as N}from"./el-col-1770158d.js";import{d as n,u as H,c as w,o as l,a as _,r as m,n as f,b as r,e as y,f as V,g as k,w as h,h as g,i as u,t as z,_ as F,j as A,k as M,F as j}from"./index-bf5dd43b.js";const D=n({name:"ElContainer"}),I=n({...D,props:{direction:{type:String}},setup(o){const e=o,t=H(),s=i("container"),a=w(()=>e.direction==="vertical"?!0:e.direction==="horizontal"?!1:t&&t.default?t.default().some(d=>{const $=d.type.name;return $==="ElHeader"||$==="ElFooter"}):!1);return(c,d)=>(l(),_("section",{class:f([r(s).b(),r(s).is("vertical",r(a))])},[m(c.$slots,"default")],2))}});var T=p(I,[["__file","/home/runner/work/element-plus/element-plus/packages/components/container/src/container.vue"]]);const R=n({name:"ElAside"}),q=n({...R,props:{width:{type:String,default:null}},setup(o){const e=o,t=i("aside"),s=w(()=>e.width?t.cssVarBlock({width:e.width}):{});return(a,c)=>(l(),_("aside",{class:f(r(t).b()),style:y(r(s))},[m(a.$slots,"default")],6))}});var E=p(q,[["__file","/home/runner/work/element-plus/element-plus/packages/components/container/src/aside.vue"]]);const G=n({name:"ElFooter"}),J=n({...G,props:{height:{type:String,default:null}},setup(o){const e=o,t=i("footer"),s=w(()=>e.height?t.cssVarBlock({height:e.height}):{});return(a,c)=>(l(),_("footer",{class:f(r(t).b()),style:y(r(s))},[m(a.$slots,"default")],6))}});var b=p(J,[["__file","/home/runner/work/element-plus/element-plus/packages/components/container/src/footer.vue"]]);const K=n({name:"ElHeader"}),L=n({...K,props:{height:{type:String,default:null}},setup(o){const e=o,t=i("header"),s=w(()=>e.height?t.cssVarBlock({height:e.height}):{});return(a,c)=>(l(),_("header",{class:f(r(t).b()),style:y(r(s))},[m(a.$slots,"default")],6))}});var x=p(L,[["__file","/home/runner/work/element-plus/element-plus/packages/components/container/src/header.vue"]]);const O=n({name:"ElMain"}),P=n({...O,setup(o){const e=i("main");return(t,s)=>(l(),_("main",{class:f(r(e).b())},[m(t.$slots,"default")],2))}});var B=p(P,[["__file","/home/runner/work/element-plus/element-plus/packages/components/container/src/main.vue"]]);C(T,{Aside:E,Footer:b,Header:x,Main:B});v(E);v(b);const Q=v(x);v(B);const U="/favicon.ico",W=u("div",null,[u("a",{href:"https://essay.zag13.com/",target:"_blank"},[u("img",{class:"h-8",src:U,alt:"essay.zag13.com"})])],-1),X=n({__name:"AppHeader",setup(o){const e=V(!1);return(t,s)=>{const a=S,c=N,d=Q;return l(),k(d,{class:"bg-slate-100 w-full flex items-center"},{default:h(()=>[g(c,{class:"w-full"},{default:h(()=>[g(a,{class:"flex p-1 justify-between"},{default:h(()=>[W,u("button",{onClick:s[0]||(s[0]=()=>e.value=!e.value)},[u("span",null,z(e.value?"To Be":"Not To Be"),1)])]),_:1})]),_:1})]),_:1})}}}),Y={};function Z(o,e){const t=A("router-view");return l(),k(t,null,{default:h(({Component:s,route:a})=>[(l(),k(M(s),{key:a.path}))]),_:1})}const ee=F(Y,[["render",Z]]),ne=n({__name:"index",setup(o){return(e,t)=>(l(),_(j,null,[g(X),g(ee)],64))}});export{ne as default};
