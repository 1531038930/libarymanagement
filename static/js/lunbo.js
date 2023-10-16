function lb(){
    const imgs=[{src:"../src/img1.png",title:"战神",bgc:"gray"},{src:"../src/img2.png",title:"战地",bgc:"blue"},{src:"../src/img3.png",title:"sky",bgc:"brown"}]
    const img=document.getElementById("imgs")
    const p=document.querySelector(".lbDown p")
    const bgc=document.querySelector(".lb")
    img.src=imgs[i].src
    p.innerHTML=imgs[i].title
    bgc.style.backgroundColor=imgs[i].bgc
    const points=document.querySelectorAll(".quickChange li")
    document.querySelector("li.active").classList.remove("active")
    points[i]["classList"].add("active")
    i++
    if (i==3){
        i=0
    }
}
let i=0
lb()
let si=setInterval(lb,3000)
let listenimg=document.getElementById("imgs")
listenimg.addEventListener('mouseenter',function (){
    clearInterval(si)
})
listenimg.addEventListener('mouseleave',function (){
    si=setInterval(lb,3000)
})
let listenpoints=document.querySelectorAll(".quickChange li")
for (let n=0;n<listenpoints.length;n++){
    listenpoints[n].addEventListener('click',function (){
        clearInterval(si)
        i=n
        lb()
        si=setInterval(lb,3000)
    })
}
let search=document.querySelector(".search")
search.addEventListener('focus',function () {
    let sl=document.getElementsByClassName("searchlist")
    // let sl=document.querySelector(".searchlist")
    // console.log(sl.length)
    sl[0]["style"]['display']="block"
})
search.addEventListener('blur',function () {
    let sl=document.getElementsByClassName("searchlist")
    // let sl=document.querySelector(".searchlist")
    // console.log(sl.length)
    sl[0]["style"]['display']="none"
})
let scroll=document.querySelector(".scroll")
scroll.addEventListener("scroll",function (e) {
    console.log(scroll.scrollTop)
})
const qs1=document.querySelector(".backtop")
qs1.addEventListener("click",function () {
    console.log(window.scrollY)
    window.scrollTo({top:0,behavior:'smooth'})
})
window.addEventListener("scroll",function (){
    const qs=document.querySelector(".backtop")
    const header=document.querySelector(".header")
    const pos=document.querySelector(".input")
    console.log(pos.offsetTop)
    if (window.scrollY>pos.offsetTop){
        qs.style.display='block'
        if(window.scrollY>pos.offsetTop+10){
            qs.style.opacity=1
        }else {
            qs.style.opacity=0
        }
        header.style.top='0px'
    }else{
        qs.style.display='none'
        header.style.top='-100px'
    }
})