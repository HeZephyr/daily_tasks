package main
import "fmt"
func main() {
	s := "<!DOCTYPE html>
<html lang="zh-CN">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,Chrome=1" />
  <meta name="viewport" content="width=device-width,initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0,user-scalable=no" />
  <meta name="format-detection" content="telephone=no" />
  <title>计算机科学与技术学院</title>
  <meta name="keywords" content="" />
  <meta name="description" content="" />
  
<link type="text/css" href="/_css/_system/system.css" rel="stylesheet"/>
<link type="text/css" href="/_upload/site/1/style/2/2.css" rel="stylesheet"/>
<link type="text/css" href="/_upload/site/00/44/68/style/128/128.css" rel="stylesheet"/>
<link type="text/css" href="/_js/_portletPlugs/sudyNavi/css/sudyNav.css" rel="stylesheet" />
<link type="text/css" href="/_js/_portletPlugs/datepicker/css/datepicker.css" rel="stylesheet" />
<link type="text/css" href="/_js/_portletPlugs/simpleNews/css/simplenews.css" rel="stylesheet" />

<meta content="A5YrmdTHjjao4RQy0zA42cKBCIchwbmr" r="m"><!--[if lt IE 9]><script r='m'>document.createElement("section")</script><![endif]--><script type="text/javascript" r='m'>$_ts=window['$_ts'];if(!$_ts)$_ts={};$_ts.nsd=21076;$_ts.cd="qhSDrpAlhk7qtGWjqrq5tG39r1LmlqVciaqrqnLqlkgjrSqmrk7qtGVjtpq5tGqjqAl7cqAIiaqXEPE4qcLltGVAHPLkxqG4qcLjtG3AHPLltGVrhq9lEu7qqAWjqrq5iaqjrPLolkg4qqlmtGEAHO7qtGqjqpq5tG3ocGAwtG3ArAW4qcLltGVAms7qhPEhiaqjqPLklkgjqAQhrcq4qcLrtG3AHPLmlqVNiaqjqnLqlkg4qcLDtGQAm1LlhrgDqs3SqGVSkA3OR3s3NQhV0fz8nQBaEdqdbNc9ZAaG3EoUNfpgPJRXXMVqrAEoruGNkZR881VTAl9aWoE7YCRnY022p0fCIvrJsbylEkTQ320fUPNOw6QVQUqBQvlLR92ChvJGQomXKbW7hPZGt6JhQCRwKcyNQ2YjhCyaMb7.Y6JBMc27QbzYtbTTFKg.JSYLwbSNUn728OZu3CNGKmfr1OREHV3Nsl2CpsVep1Ve3bfCMD9XFUxYFPeXQDy.tTwuMCaNMox.Y1NzwK2XtuQYFoR7FYZ.xUTvWkxBVsyLqlpEWVg0YOeaMKTjYUW2HKxLwbSNhb2nYKLBMvrXFnzUQCe7tKSnFTVjMUYNMP76mKfCMD2ItnYL1KYEMompMmxvJsG.wmTV3DxbwmwOxsTPFoR7F1yNQ2YjhCyaMb7.Y6JBMc27QbzYtbTTFKg.JSYLwbSNUn72Fbfx8vR0F6mOWD2TYOYOYUfNibTFp1Ve3bfCMD9XFUxYFPeXQDy.tTwuMCaNMox.Y1NzwK2XtuQYFoR7FYZ.xKJ5R9x78ueBVKpxhPzuwKSZH2x_JVV2HKxLwbSNhb2nYKLBMvrXFnzUQCe7tKSnFTVjMUYNMP76mKfCMD2ItnY3MuN9Q92iFCeNsO2GMmVC3DeqKsSmxsTPFoR7F1yNQ2YjhCyaMb7.Y6JBMc27QbzYtbTTFKg.JSYLwbSNUn72VOxzUCmSQsqtHvpsJUf81DyEss2FAnVe3bfCMD9XFUxYFPeXQDy.tTwuMCaNMox.Y1NzwK2XtuQYFoR7FYZ.xYwEF2RXMKrPKCyKM6SnM2YFIDeSF9W2HKxLwbSNhb2nYKLBMvrXFnzUQCe7tKSnFTVjMUYNMP76mKfCMD2ItnYUYKpzA9rIsKrP1KeDW0eXw6rcJKxOxsTPFoR7F1yNQ2YjhCyaMb7.Y6JBMc27QbzYtbTTFKg.JSYLwbSNU6L.32TzwK2XtC2WFbZ7xuE2tS75WnENEm25A6V.3bf53D7AtDTTRDWN32TuhKJ.EczNsDNBhkVXFUxYF2g7UbSnFTVjMUYNMP72c1ABEO0nx17MKcz0FDN73Vz_wC29Moz0VDYBIoRf86p3FDrLQDJN3lY68UST8veupUR6IoYaMb2J8vJLIomXF9YCF62G8vw4YbfbIDNL8CfwwUeBRoY6wlYLworyMoxbKoe7wUmj8C2t36eXMKe5hfwLwbm9hDxOscfGMPQXFUxYFPZatKSnFTYMhTz.wCz.mKfCMD9XxsQMtPlgJ1q.xKQX3DmGRKz4sCNP3v2T3byrMC0NtCSnR2GLRbpahDpNDnNzwK2XtuQYFoR7FYZ.UmwuMCaNMox.Y1LfWcQXxsZxxcZPKnNuF2mZ3USSwbp4sDwXIoJ.wUJWRCeuMDe4QYRPQDY6wDxTpUrSRKz2Mb2Q8vpLwUyaM2wPw6Y.MDx.pUmuhbef86rrMUNgQUyaIYeSI6G6Ive9pUyGRbJ5ICrmIbyvF17NFVSv3PyGRvENKbGfhDz6FC7YJPeXQDy.KfYwMUYNMPzNsDNBhcQnx17MHOWPt1rRhTzZFCzfIbSuAKVjhbz6Rb3U3CwTtCmNxpYLwbSNhOQ.Y6JBMmGXUD2WFbZ7F6w.FSV9WnENEkZ6DnLfUPT98Cm3RK3NtCSnR2GLRbpahDpNDnNzwK2XtuQYFoR7FYZ.UmwuMCaNMox.Y1LfWcQXxsZxxcZPKnNPFmY_FCmZwbx0s1ZzMvrORP2Q3v3XRba2tTwuMCaNWPzNsDNBU12RFUxYFPeXQDy.tS7CE1aviOQ2m1pwtoe9wCNQwUe2MDw4F9rP3vp9IbNPYKwfIDw.FDJKwKNz3oRTI9YGRbm9RDJ0pKySIDxzFKrtMbLNtCSnR2GLRbpahDpNDnNzwK2XtuQYFoR7FYZ.UmwuMCaNMox.Y1LfWcQXxsZxxcZPKnNe3TNTI6JB3KT4sKfC3Kp5MK2WRCxZICRdRpqLMUY93n2PUUWz3C7GtC2WFbZ7J1yNQ2YjUPS8Mox.Y1NzwK2XtnVxxcZPHsE2tSz3tCm7QDpTKDr6hKpSwDxVRDey8KJOF2rTQD2jMop4KDYXIDpf3CNmRCxBtDwbwrrLwbm9hDxOscfGMPQXFUxYFPZatKSnFTYMhTz.wCz.mKfCMD9XxsQMtPlgJ1q.xKwTw6VzRbxNsULjhbz6Rb3U3CwTtCmNxpYLwbSNhOQ.Y6JBMmGXUD2WFbZ7F6w.FSV9WnENEkZ6DnLfUPT4I6zJQ1LXF6wbRfwOR6V.3K92mKfCMD9XJnzUQCe7KcyQFVSjMnS.wCz.m1AuEP9GHuQMtPm8hKNLQ9NOICSeIvp51vmb3v26wUQEhnTXQDJbtYSPwPybMcV.Y6JBMc9ntC2WFbewtY2NQ2YjhCyaMb7.Du3fhcQ_JnVYxm0Lwvy0MTf0RP9.MoxbKcfbRoq7RD9MtbTTFKg.JSYLwbSNUnzQY6JBMc27QbzYtPlaxng2Ha99hnxMtbrepKm_Rv22I6z8QCxXFKz43lzvR6Y9wbx4UbpGQo29IvzIQ6pT8KpuI0T_RKmbwCzuAkpdRbpfQbzJI6ebRDmnF0TdhbY.IbxOUCzuIDrGMbzY8bw_IUNGRlYP86reIbJuKbYPIDm.woAAtDTTRDWN32TuhKJ.EczNsDNBhkVXFUxYF2g7UbSnFTVjMUYNMP72c1ABEO0nx17MKcz_3Km08TJLwbm9IvJNsDwPIoRLRDTmIDwgRn7NFVSv3PyGRvENKbGfhDz6FC7YJPeXQDy.KfYwMUYNMPzNsDNBhcQnx17MHOWPt1rRhTJ2MC2vwKxXAKezIDefFCyM8b2OM6N4MYNBQCmCQUzzKbpa8vm5MKpmM6r28KzT8TrNw6JbIbyLsspdFKNCJvzAFUwZMDS08Tr.QoS7MbR2pKePM62jIbziwvrZMURLR9Yf8KxeQvWXmCfC3KA73bJWtDrXxnyNQ2YjhuqNMox.YY7BKCz6FC7YFoR7F1g2JS7jEsGTEc72RnyyQ6eB86TtwKRnQCz939YZQ6SdFUeupUS.8Dx5IKe8Mve_3vye3Yz.I6NfIvynVbJ63Kz2I6pmI6w2RDe4I9NT3USZFUWNACxd8oSjICpr8vzyMKNu80mTRo2y8vpGlcfzwKpOtKxwQnTOFnq.FVSjMnaThb2nYKN8hmy7QbzYtbTTFKg.xA99hnEeWPV.DTG.QDT4wDyFwUe2FKw.R27BhKya3C3NUDxChbR7x1zUQCe7tsE.FVSjMT0NKK2nYKLBMvrXFn7MJPl7xuL6xpV9U1fL3KN2YoNXRK2ORUyqhnTXQDJbtYSPwPybMcV.Y6JBMc9ntC2WFbewtY2NQ2YjhCyaMb7.Du3fhcQ_JnVYxm0LMKwT80JCMP9.MoxbKcfbRoq7RD9MtbTTFKg.JSYLwbSNUnzQY6JBMc27QbzYtPlaxng2Ha99hnE.tc2NsDwPhbrbQP2FFcl7F6w.FSVChCyaMbzFmYTzwK2XtC2WFbZ7xuE2tS75WnENEm9B1vRP3CqB3vexFbSnRoRXFmw_wbmuQUrb1Czd3KraFDxU3KNX3ce.R2z_F9S93cNdwVwqYVrD8CRQwDS2wUm4QYr0w6Y9MvpXYcZzMvrORP2Q3v3XRba2tTwuMCaNWPzNsDNBU12RFUxYFPeXQDy.tS7CE1aviOQ2m1pwtDJ.FKJmFvra8KefM2wPtPy.wCRbmCJvw1zfF1VYFoR7F1g6tTwuMCS3h2SNsDNBhDz6FC7YxkWPt1qdJS7jEYgzFvw9KCNNIoRfhP2UQCxftCwOQfw2M1ENMox.Y1LuhDz6FCz9t2fXQDy.tTwuMCaNEkQ2m1A4WcQXxY2xR6wO3DSu397XhKJLRCx4lcfzwKpOtKxwQnTOFnq.FVSjMnaThb2nYKN8hmy7QbzYtbTTFKg.xA99hnEeWPV.DTG.3K2SQ6xmIvrvF17NMYLLMUY93n2PUUWz3C7GtC2WFbZ7J1yNQ2YjUPS8Mox.Y1NzwK2XtnVxxcZPHsE2tSz3tCyzFvR4YCyGRU273vJmFCR08KeP39TBMKwyMby2sKJawo2XMbYhFKSz8KyXRVzOQuxyMby21CJuID2.RKr3QbRZFKz28TELtUSB8br.m6YGFKr7RDxmMozv8KfLRVY7RoJyFDWCcs3jhbz6Rb3U3CwTtCmNxpYLwbSNhOQ.Y6JBMmGXUD2WFbZ7F6w.FSV9WnENEkZ6DnLfUPTfQCmmRUqzIKR08TNThbSnQvp4KvRd3Cfv3vyxI6eOMvN2RYp_3Dwy3KxjY1zTQv2f3byQFKYbMDmP8TNOICJb3CRL16NG3vR5RKfKMKrnFvybwlT0McTj36p4KowyQ6T_8CREMDrZRDynF2wBwoG7hD2nKDQzRKm6tKpUxceXQDy.ta9jMUYNM2NFRqack2T.HCmFwlmiHOE6WvSSwkVuiUppQDqaR0mhHDTdYb96HYpEHmfW1lyYwKNtq9JApKQnMYYkFTQnAO2hsbyfFTxAV9VuFl0gKUpOQ0RUMVmuQbrLHDNtJY9S8DmqMDJvpn7uHOyF3kxPtCf0HmR5K2xDRTVCVCS9sn7uHOyF3kxPtCf0HoQ5V2pj3CN9MDpUA17uHOyF3kxPtCf0HlyKw9zHYvw2FDePi1zLwkSIssmCWUf53mNnAYL530pIkCQuiKEnVlYAHUJDw9fqsaSnFCGuJkSzI2JlwlwDQ0zF3TlXJk75RmZ631yBQsSLkswbW9msAUpC3CaXJk75RmZ631yBQsaNckZ_3CWuR12swAWqrAW6ruG6qSwzwvJ6q6GztPz7QbR3tDR9QcS9FAW5WaqkWug5Kbqnqq3uWsVykkEaWO3aHG9nJuqmrsZ0qkA0Wklcmgys8yia9s_QBFtBAC3Gn00otY8n_7Pa3HegHaVkrAQmqkgDEdd58FypeTPj53hedpbwyoDjkJ5xyed.6UamInCE0S2XqqZSJuEqJO30JsqaruG6qaL0JsV5cuWSWO30rnSBwO7gQ2fYVbeJUlR7hkR1IlpEw1z2IURqRkxxMUYkH6T2RKgnM6EaVUE_qqVcraQgJGWrWGHgWlWj3bmbhbRj1nNPQCGXRvmAtbmfMPy2R90j3KTThbYXYnNGhDROR1zFwUq7RvmStTNNMPSbQvQ.UKw6hDmXFnzwFKq73KTutTTN3oGNR6ePmKJ7wc29FC3Y3DeOtKpLFpYGMoANRDfNs1N9RDQX3DyItbp_tKYBFrYXM6qNFvw6mKy63n2.RDGYMCTNtKzN3pYBMUWNF6pLmKSahDeOw1zswbA7M6xP3rY7QDe.hbTbsCQBMDSGtCzK31ezRoE.F2rPhCTTJPzj1KQBMsmOtCNxJ1ezwKV.F2xahCTjRPzjYvqBMKSBtCNxwPezwb7.F92ORcSj3D0.YbY4hDN23czKFbl7FbyPtTRnR1ST8nz6ADVBQDpSt6w8FneawKZ.w02Th6rTwPz6UKqBQoYCInzxFDA7wKz9t02jFPSTQKW.1De2hopCInzhRnenFcyTR2AjQKTu3czTUDlBQbrSt6YJ3cenRoq.wm3jQDxNhvpfA1NawCAXworItvrC3Py0I9LjwCm6hvJbAnN6FKWXQCyI3ce0FvrT32wnh6pewPzusKqBwvRbt6mtw1eCwvV.QVmPh6wjRDl.sb22ROQXQDNHtvpzRcyaR9ZjworChve2YcN4QDGXICzttvz7wvL.I0Njh6fLMPzemUSjQc2OwCaYwbeNwv3.M92vhCS93qQoiAxQxmN3raJbE2ScrAR3ETeXqapFx2friAQaqqQTJaQgJaWrJqEoru9";if($_ts.lcd)$_ts.lcd();</script><script type="text/javascript" charset="utf-8" src="/LIWghRUPB4QK/oxYRCeR1jjgO.199cf1b.js" r='m'></script><script language="javascript" src="/_js/sudy-jquery-autoload.js" jquery-src="/_js/jquery-1.9.1.min.js" sudy-wp-context="" sudy-wp-siteId="68"></script>
<script language="javascript" src="/_js/jquery-migrate.min.js"></script>
<script language="javascript" src="/_js/jquery.sudy.wp.visitcount.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/sudyNavi/jquery.sudyNav.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/datepicker/js/jquery.datepicker.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/datepicker/js/datepicker_lang_HK.js"></script>
<link href="/_upload/tpl/04/02/1026/template1026/css/base.css" rel="stylesheet" type="text/css">
  <link href="/_upload/tpl/04/02/1026/template1026/css/11ml.css" rel="stylesheet" type="text/css">
  
  <script type="text/javascript" src="/_upload/tpl/04/02/1026/template1026/js/jQuery.meanMenu.js"></script>
  <script type="text/javascript" src="/_upload/tpl/04/02/1026/template1026/js/jquery.bxslider.min.js"></script>
  <link href="/_upload/tpl/04/02/1026/template1026/extends/extends.css" rel="stylesheet">
  <script type="text/javascript" src="/_upload/tpl/04/02/1026/template1026/extends/extends.js"></script>
  <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
  <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
  <!--[if lt IE 9]>
  <script src="https://cdn.bootcss.com/html5shiv/3.7.3/html5shiv.min.js"></script>
  <script src="https://cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
<![endif]-->
  <link rel="stylesheet" href="/_upload/tpl/04/02/1026/template1026/css/owl.carousel.min.css">
  <link rel="stylesheet" href="/_upload/tpl/04/02/1026/template1026/css/owl.theme.default.css">
  <script src="/_upload/tpl/04/02/1026/template1026/js/owl.carousel.min.js"></script>
</head>

<body>
  <div class="top-se">
    <div class="container">
      <div class="t_se clearfix" frag="窗口0" portletmode="search">            <form method="post" action="/_web/_search/api/search/new.rst?locale=zh_CN&request_locale=zh_CN&_p=YXM9NjgmdD0xMDI2JmQ9NDE2MCZwPTEmbT1TUyZjPXcwJnY9QyY_" target="_blank">
                <input name="keyword" type="text" class="se_txt">
                <input name="submit" type="submit" class="se_sub" value="搜 索">
                <a class="se_close" href="javascript:;"></a>
            </form>
        </div>
    </div>
  </div>
  <div class="top">
    <div class="container">
      <div class="row">
        <div class="col xs-12 sm-12 md-4 lg-5">
          <div class="logo">
            <div class="logo_table">
              <div class="logo_cell"> <a href="/main.htm"><img src="/_upload/tpl/04/02/1026/template1026/images/logo.png"></a> </div>
            </div>
          </div>
        </div>
        <div class="col xs-12 sm-12 md-8 lg-7">
          <div class="top_r">
            <div class="top_r_a"><a class="a1" href="http://www.nuaa.edu.cn/" target="_blank">南航主页</a><a class="a2" href="javascript:void(0)">ENGLISH</a><a class="a3" href="javascript:;"></a></div>
            <div class="top_nav hidden-xs" frag="窗口1" portletmode="simpleSudyNavi" configs="{'c1':'1','c7':'2','c4':'_self','c3':'30','c6':'0','c8':'2','c9':'0','c2':'1','c5':'2'}" contents="{'c2':'0', 'c1':'/学院概况,/师资队伍,/科学研究,/教学工作,/学生工作,/党群工作,/工会与校友工作,/常用下载'}">
              
              <ul class="clearfix pc_menuCon">
                <li class="active"><a href="/main.htm">首页</a></li>
                
                <li><a href="/1954/list.htm" target="_self">学院概况</a>
                  
                  <ul>
                    
                    <li><a href="/xyjj/list.htm" target="_self">学院简介</a></li>
                    
                    <li><a href="/1965/list.htm" target="_self">Introduction of the College</a></li>
                    
                    <li><a href="/1966/list.htm" target="_self">领导班子</a></li>
                    
                    <li><a href="/1967/list.htm" target="_self">学院概况</a></li>
                    
                    <li><a href="/1968/list.htm" target="_self">历史沿革</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1956/list.htm" target="_self">师资队伍</a>
                  
                  <ul>
                    
                    <li><a href="http://graduate.nuaa.edu.cn/gmis5/dsfc/dsfc_yx_new" target="_self">硕导博导信息</a></li>
                    
                    <li><a href="/7382/list.htm" target="_self">学院教师</a></li>
                    
                    <li><a href="/7383/list.htm" target="_self">办公室人员</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1957/list.htm" target="_self">科学研究</a>
                  
                  <ul>
                    
                    <li><a href="/1975/list.htm" target="_self">科研动态</a></li>
                    
                    <li><a href="/1976/list.htm" target="_self">表格下载</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1958/list.htm" target="_self">教学工作</a>
                  
                  <ul>
                    
                    <li><a href="/zyjs/list.htm" target="_self">专业介绍</a></li>
                    
                    <li><a href="/1977/list.htm" target="_self">教学动态</a></li>
                    
                    <li><a href="/jxzd/list.htm" target="_self">教学制度</a></li>
                    
                    <li><a href="/1978/list.htm" target="_self">本科生教学</a></li>
                    
                    <li><a href="/1979/list.htm" target="_self">研究生教学</a></li>
                    
                    <li><a href="/1980/list.htm" target="_self">教学改革与成果</a></li>
                    
                    <li><a href="/1981/list.htm" target="_self">资料下载</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1959/list.htm" target="_self">学生工作</a>
                  
                  <ul>
                    
                    <li><a href="/1982/list.htm" target="_self">学生活动</a></li>
                    
                    <li><a href="/1993/list.htm" target="_self">招生信息</a></li>
                    
                    <li><a href="/1983/list.htm" target="_self">研究生会</a></li>
                    
                    <li><a href="/1994/list.htm" target="_self">就业信息</a></li>
                    
                    <li><a href="/1984/list.htm" target="_self">共青团</a></li>
                    
                    <li><a href="/1985/list.htm" target="_self">学生党建</a></li>
                    
                    <li><a href="/1986/list.htm" target="_self">获奖情况</a></li>
                    
                    <li><a href="/1987/list.htm" target="_self">通知公告</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1960/list.htm" target="_self">党群工作</a>
                  
                  <ul>
                    
                    <li><a href="/qgdjgzybzb/list.htm" target="_self">全国党建工作样板支部</a></li>
                    
                    <li><a href="/1988/list.htm" target="_self">党员管理</a></li>
                    
                    <li><a href="/1989/list.htm" target="_self">主题活动</a></li>
                    
                    <li><a href="/1990/list.htm" target="_self">组织建设</a></li>
                    
                    <li><a href="/1991/list.htm" target="_self">党员发展</a></li>
                    
                    <li><a href="/1992/list.htm" target="_self">党校工作</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1962/list.htm" target="_self">工会与校友工作</a>
                  
                  <ul>
                    
                    <li><a href="/7384/list.htm" target="_self">教师工作指南</a></li>
                    
                    <li><a href="/tmhxqgzshzn/list.htm" target="_self">天目湖校区工作生活指南</a></li>
                    
                    <li><a href="/16162/list.htm" target="_self">教职工之家</a></li>
                    
                    <li><a href="/16163/list.htm" target="_self">校友返校指南</a></li>
                    
                    <li><a href="/16164/list.htm" target="_self">基金工作指南</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/cyxz/list.htm" target="_self">常用下载</a>
                  
                  <ul>
                    
                    <li><a href="/aqgz/list.htm" target="_self"> 安全工作</a></li>
                    
                    <li><a href="/rsxz/list.htm" target="_self">人事工作</a></li>
                    
                    <li><a href="/bkjx/list.htm" target="_self">本科生培养</a></li>
                    
                    <li><a href="/yjspy/list.htm" target="_self">研究生培养</a></li>
                    
                    <li><a href="/kxyj/list.htm" target="_self">科学研究</a></li>
                    
                    <li><a href="/xsgz/list.htm" target="_self">学生工作</a></li>
                    
                    <li><a href="/djqt/list.htm" target="_self">党建群团</a></li>
                    
                    <li><a href="/xygz/list.htm" target="_self">校友工作</a></li>
                    
                    <li><a href="/xyLogo/list.htm" target="_self">学院Logo</a></li>
                    
                  </ul>
                  
                </li>
                
              </ul>
              
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="phone_menu_t visible-xs">
    <div class="phone_menu" style="display: none" frag="窗口01" portletmode="simpleSudyNavi" configs="{'c1':'1','c7':'2','c4':'_self','c3':'30','c6':'0','c8':'2','c9':'0','c2':'1','c5':'2'}" contents="{'c2':'0', 'c1':'/学院概况,/师资队伍,/科学研究,/教学工作,/学生工作,/党群工作,/工会与校友工作,/常用下载'}">
      
      <ul>
        
        <li><a href="/1954/list.htm" target="_self">学院概况</a>
          
          <ul>
            
            <li><a href="/xyjj/list.htm" target="_self">学院简介</a></li>
            
            <li><a href="/1965/list.htm" target="_self">Introduction of the College</a></li>
            
            <li><a href="/1966/list.htm" target="_self">领导班子</a></li>
            
            <li><a href="/1967/list.htm" target="_self">学院概况</a></li>
            
            <li><a href="/1968/list.htm" target="_self">历史沿革</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1956/list.htm" target="_self">师资队伍</a>
          
          <ul>
            
            <li><a href="http://graduate.nuaa.edu.cn/gmis5/dsfc/dsfc_yx_new" target="_self">硕导博导信息</a></li>
            
            <li><a href="/7382/list.htm" target="_self">学院教师</a></li>
            
            <li><a href="/7383/list.htm" target="_self">办公室人员</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1957/list.htm" target="_self">科学研究</a>
          
          <ul>
            
            <li><a href="/1975/list.htm" target="_self">科研动态</a></li>
            
            <li><a href="/1976/list.htm" target="_self">表格下载</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1958/list.htm" target="_self">教学工作</a>
          
          <ul>
            
            <li><a href="/zyjs/list.htm" target="_self">专业介绍</a></li>
            
            <li><a href="/1977/list.htm" target="_self">教学动态</a></li>
            
            <li><a href="/jxzd/list.htm" target="_self">教学制度</a></li>
            
            <li><a href="/1978/list.htm" target="_self">本科生教学</a></li>
            
            <li><a href="/1979/list.htm" target="_self">研究生教学</a></li>
            
            <li><a href="/1980/list.htm" target="_self">教学改革与成果</a></li>
            
            <li><a href="/1981/list.htm" target="_self">资料下载</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1959/list.htm" target="_self">学生工作</a>
          
          <ul>
            
            <li><a href="/1982/list.htm" target="_self">学生活动</a></li>
            
            <li><a href="/1993/list.htm" target="_self">招生信息</a></li>
            
            <li><a href="/1983/list.htm" target="_self">研究生会</a></li>
            
            <li><a href="/1994/list.htm" target="_self">就业信息</a></li>
            
            <li><a href="/1984/list.htm" target="_self">共青团</a></li>
            
            <li><a href="/1985/list.htm" target="_self">学生党建</a></li>
            
            <li><a href="/1986/list.htm" target="_self">获奖情况</a></li>
            
            <li><a href="/1987/list.htm" target="_self">通知公告</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1960/list.htm" target="_self">党群工作</a>
          
          <ul>
            
            <li><a href="/qgdjgzybzb/list.htm" target="_self">全国党建工作样板支部</a></li>
            
            <li><a href="/1988/list.htm" target="_self">党员管理</a></li>
            
            <li><a href="/1989/list.htm" target="_self">主题活动</a></li>
            
            <li><a href="/1990/list.htm" target="_self">组织建设</a></li>
            
            <li><a href="/1991/list.htm" target="_self">党员发展</a></li>
            
            <li><a href="/1992/list.htm" target="_self">党校工作</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1962/list.htm" target="_self">工会与校友工作</a>
          
          <ul>
            
            <li><a href="/7384/list.htm" target="_self">教师工作指南</a></li>
            
            <li><a href="/tmhxqgzshzn/list.htm" target="_self">天目湖校区工作生活指南</a></li>
            
            <li><a href="/16162/list.htm" target="_self">教职工之家</a></li>
            
            <li><a href="/16163/list.htm" target="_self">校友返校指南</a></li>
            
            <li><a href="/16164/list.htm" target="_self">基金工作指南</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/cyxz/list.htm" target="_self">常用下载</a>
          
          <ul>
            
            <li><a href="/aqgz/list.htm" target="_self"> 安全工作</a></li>
            
            <li><a href="/rsxz/list.htm" target="_self">人事工作</a></li>
            
            <li><a href="/bkjx/list.htm" target="_self">本科生培养</a></li>
            
            <li><a href="/yjspy/list.htm" target="_self">研究生培养</a></li>
            
            <li><a href="/kxyj/list.htm" target="_self">科学研究</a></li>
            
            <li><a href="/xsgz/list.htm" target="_self">学生工作</a></li>
            
            <li><a href="/djqt/list.htm" target="_self">党建群团</a></li>
            
            <li><a href="/xygz/list.htm" target="_self">校友工作</a></li>
            
            <li><a href="/xyLogo/list.htm" target="_self">学院Logo</a></li>
            
          </ul>
          
        </li>
        
      </ul>
      
    </div>
  </div>
  <script type="text/javascript">
  $(function() {
    $('.pc_menuCon li').hover(function() {
      $('ul', this).slideDown(200);
    }, function() {
      $('ul', this).slideUp(100);
    });
    $('.a3').click(function() {
      $('.top-se').stop().slideToggle();
    });
    $('.se_close').click(function() {
      $('.top-se').stop().slideUp(200);;
    });
  });
  jQuery('.phone_menu').meanmenu();
  $(window).scroll(function() {
    var scrollHeight = $(document).scrollTop();
    if (scrollHeight > 340) {
      $('.phone_menu_t').addClass('lighted-fixed');
    } else {
      $('.phone_menu_t').removeClass('lighted-fixed');
    }
  });
  </script>
  <div class="banner hidden-xs" frag="窗口2" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'6','c2':'图片路径','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'15','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'1', 'c1':'/PC首页幻灯'}"><div id="wp_news_w2"> 

    <ul class="bxslider">
      
      <li><img src="/_upload/article/images/01/04/523d814246dda4c114c9425dce30/fd8e6513-4264-4a38-9b99-567e8a804b4c.png" /></li>
      
      <li><img src="/_upload/article/images/2c/68/7bd85cd148e287ce7296d350a07f/dd5534aa-c058-43bc-86f3-b03d80406e03.png" /></li>
      
      <li><img src="/_upload/article/images/de/1f/d41dead342b98bcd2fc4628e1067/681fed19-48c5-4eed-9828-01b0a69e130d.jpg" /></li>
      
      <li><img src="/_upload/article/images/86/88/f2ef68b44c5e9a69ddd5b32f819f/bbfa325d-5e6c-4e38-8d5b-531597e40432.png" /></li>
      
      <li><img src="/_upload/article/images/d4/9c/2ba2b6a34b1d9fc89fcc925c3bb7/00fab851-97ca-4ec2-b88a-f3d3c6a84d78.png" /></li>
      
      <li><img src="/_upload/article/images/45/45/a46e74f24558b84ce2c9d7e02af5/24004537-f4bf-40aa-83b1-0f9e7afcb15e.png" /></li>
      
    </ul>
  </div> 
</div>
  <script type="text/javascript">
  $('.bxslider').bxSlider({
    mode: 'fade',
    auto: true,
    autoControls: true,
    autoHover: true,
    controls: false
  });
  </script>
  <!-- pc-banner -->
  <div class="banner visible-xs" frag="窗口02" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'6','c2':'图片路径','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'15','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'1', 'c1':'/手机首页幻灯'}"><div id="wp_news_w02"> 

    <ul class="bxslider02">
      
      <li><img src="/_upload/article/images/20/35/0e7ca14f46ea844bd475df8ebf37/48f77e80-bf7c-474d-a504-5ea301c1b032.png" /></li>
      
      <li><img src="/_upload/article/images/74/5b/9dda4c0446fd93b0326d50a71b45/91732d13-5f6e-4d78-96ef-7b4c3c2f565b.png" /></li>
      
      <li><img src="/_upload/article/images/21/54/835e7c8d47a3ac33755a44a53147/a96926c2-fd26-4c54-b1e5-cace00df4e3c.jpg" /></li>
      
      <li><img src="/_upload/article/images/70/99/35cbc4664cb48bdf524a41fac7a3/669f44d5-c735-4e4b-a6a4-d306f826e255.png" /></li>
      
      <li><img src="/_upload/article/images/08/98/6c53849b4957a1216a245e5c91d4/8a16e31c-8887-4390-ae2c-807a4c763fc0.png" /></li>
      
      <li><img src="/_upload/article/images/2e/c2/c5683ded49faafa8ca32ba4bb670/5a4d4182-bc67-4584-904d-a7e6ced069f6.jpg" /></li>
      
    </ul>
  </div> 
</div>
  <script type="text/javascript">
  $('.bxslider02').bxSlider({
    mode: 'horizontal',
    auto: true,
    autoControls: true,
    autoHover: true,
    controls: false
  });
  </script>
  <!-- phone-banner -->
  <div class="type">
    <div class="container">
      <div class="row">
        <div class="col xs-4 sm-4 md-2 lg-2">
          <div class="type-a type-p"> <a href="/xyjj/list.htm">
              <p>学院简介</p>
            </a> <span class="an"></span> </div>
        </div>
        <div class="col xs-4 sm-4 md-2 lg-2">
          <div class="type-b type-p"> <a href="/10850/list.htm">
              <p>本科生</p>
            </a> <span class="an"></span> </div>
        </div>
        <div class="col xs-4 sm-4 md-2 lg-2">
          <div class="type-c type-p"> <a href="/10851/list.htm">
              <p>研究生</p>
            </a> <span class="an"></span> </div>
        </div>
        <div class="col xs-4 sm-4 md-2 lg-2 col-d">
          <div class="type-d type-p"> <a href="http://bsh.nuaa.edu.cn/">
              <p>博士后</p>
            </a> <span class="an"></span> </div>
        </div>
        <div class="col xs-4 sm-4 md-2 lg-2 col-e">
          <div class="type-e type-p"> <a href="_s68/2017/0928/c11649a148148/page.psp">
              <p>校友与基金工作</p>
            </a> <span class="an"></span> </div>
        </div>
        <div class="col xs-4 sm-4 md-2 lg-2 col-f">
          <div class="type-f type-p"> <a href="/2017/1115/c11649a177680/page.htm" target="_blank">
              <p>人才引进</p>
            </a> <span class="an"></span> </div>
        </div>
        <!--<div class="col xs-6 sm-4 md-2 lg-2">
                <div class="type-a type-p"> <a href="javascript:void(0)">
                    <p>UG</p>
                    <p><i>本科</i></p>
                    </a> <span class="an"></span> </div>
            </div>
            <div class="col xs-6 sm-4 md-2 lg-2">
                <div class="type-b type-p"> <a href="javascript:void(0)">
                    <p>PG</p>
                    <p><i>研究生</i></p>
                    </a> <span class="an"></span> </div>
            </div>
            <div class="col xs-6 sm-4 md-2 lg-2">
                <div class="type-c type-p"> <a href="javascript:void(0)">
                    <p>MBA</p>
                    <p><i>工商管理硕士</i></p>
                    </a> <span class="an"></span> </div>
            </div>
            <div class="col xs-6 sm-4 md-2 lg-2 col-d">
                <div class="type-d type-p"> <a href="javascript:void(0)">
                    <p>M.Eng </p>
                    <p><i>工程硕士</i></p>
                    </a> <span class="an"></span> </div>
            </div>
            <div class="col xs-6 sm-4 md-2 lg-2 col-e">
                <div class="type-e type-p"> <a href="javascript:void(0)">
                    <p>MPAcc</p>
                    <p><i>专业会计硕士</i></p>
                    </a> <span class="an"></span> </div>
            </div>
            <div class="col xs-6 sm-4 md-2 lg-2 col-f">
                <div class="type-f type-p"> <a href="javascript:void(0)">
                    <p>EDP</p>
                    <p><i>高级经理人发展课程</i></p>
                    </a> <span class="an"></span> </div>
            </div>-->
      </div>
    </div>
  </div>
  <div class="section">
    <div class="container container-pd">
      <div class="row">
        <div class="col xs-12 sm-8 md-8 lg-8">
          <div class="section-l">
            <div class="in_title in_title_b"><span>热点新闻</span><a href="/10846/list.htm">更多>></a></div>
            <div class="section-l-con">
              <div class="row">
                <div class="col xs-12 sm-6 md-6 lg-6">
                  <div class="section-column">
                    <div class="list-img" frag="窗口3" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'7','c2':'图片路径,标题','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'30','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'1', 'c1':'/热点图片'}"><div id="wp_news_w3"> 

                      <ul class="bxslider03 bxslider03-pic">
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/ab/0e/6ed38c8149558f9ecff9874c4d01/19012e63-b7a7-4ec9-ba52-250bf0f10d7b.jpg);"></div>
                          <p><a href='/2024/1015/c11624a355830/page.htm' target='_blank' title='计算机学院举办重阳节离退休教师茶话会'>计算机学院举办重阳节离退休教师茶话会</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/c3/0b/09a3006647a0a0e2dbc6c0b6a6be/0d992420-015c-463d-a885-60d79a3b733f.png);"></div>
                          <p><a href='/2024/1016/c11624a355848/page.htm' target='_blank' title='喜迎国庆 | 青年学子心向党，知国爱国情更浓'>喜迎国庆 | 青年学子心向党，知国爱国情更浓</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/fa/83/4f5fedf34827afd995bb95059f94/9ea86eda-9308-40fc-9f9a-3dd51396661e.png);"></div>
                          <p><a href='/2024/1016/c11624a355842/page.htm' target='_blank' title='喜迎国庆 | 七十五载谱华章，强身健体筑强国'>喜迎国庆 | 七十五载谱华章，强身健体筑强国</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/eb/7a/00f3903449ea839d397beaef0a01/8f99d562-2865-4e9f-b7b5-7d26eed2f170.jpg);"></div>
                          <p><a href='/2024/0918/c11624a353696/page.htm' target='_blank' title='西安电子科技大学计算机科学与技术学院来我院调研交流'>西安电子科技大学计算机科学与技术学院来我院调研交流</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/11/37/c1f6a51c4a39be9ed36445def052/247564a3-a3a4-4146-a40f-b1939b35691a.png);"></div>
                          <p><a href='/2024/0903/c11624a352365/page.htm' target='_blank' title='计遇·智算未来 | 学院2024级新生迎新工作圆满完成'>计遇·智算未来 | 学院2024级新生迎新工作圆满完成</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/96/93/313cc7ea498e94224e738e21a160/55d0b3fa-fb5d-4459-bb82-8cb6f9549557.jpg);"></div>
                          <p><a href='/2024/0827/c11624a351892/page.htm' target='_blank' title='第十三届“中国软件杯”大学生软件设计大赛全国总决赛在我校成功举办'>第十三届“中国软件杯”大学生软件设计大赛全国总决赛在我校成功...</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/7c/20/14dc95fa4f8cbbea4d5aecf8bfdf/9a7c1f14-bc09-4207-a673-d64a6778c454.png);"></div>
                          <p><a href='/2024/0819/c11624a351709/page.htm' target='_blank' title='学院赴中国矿业大学调研交流'>学院赴中国矿业大学调研交流</a></p>
                        </li>
                        
                      </ul>
                    </div> 
</div>
                  </div>
                </div>
                <div class="col xs-12 sm-6 md-6 lg-6">
                  <div class="section-column">
                    <div class="list-ul" frag="窗口4" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'7','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/热点新闻'}"><div id="wp_news_w4"> 

                      <ul class="list-ul-li">
                        
                        <li><a href='/2024/1015/c10846a355830/page.htm' target='_blank' title='计算机学院举办重阳节离退休教师茶话会'>计算机学院举办重阳节离退休教师茶话会</a><i>10-15</i></li>
                        
                        <li><a href='/2024/1016/c10846a355848/page.htm' target='_blank' title='喜迎国庆 | 青年学子心向党，知国爱国情更浓'>喜迎国庆 | 青年学子心向党，知国爱国情更...</a><i>10-04</i></li>
                        
                        <li><a href='/2024/1016/c10846a355842/page.htm' target='_blank' title='喜迎国庆 | 七十五载谱华章，强身健体筑强国'>喜迎国庆 | 七十五载谱华章，强身健体筑强...</a><i>10-03</i></li>
                        
                        <li><a href='/2024/0930/c10846a354945/page.htm' target='_blank' title='计算机学院召开2024年下半年教职工大会'>计算机学院召开2024年下半年教职工大会</a><i>10-02</i></li>
                        
                        <li><a href='/2024/0924/c10846a354267/page.htm' target='_blank' title='13位教师入选2024年全球前2%顶尖科学家榜单'>13位教师入选2024年全球前2%顶尖科学家榜单</a><i>09-24</i></li>
                        
                        <li><a href='/2024/0918/c10846a353696/page.htm' target='_blank' title='西安电子科技大学计算机科学与技术学院来我院调研交流'>西安电子科技大学计算机科学与技术学院来我...</a><i>09-18</i></li>
                        
                        <li><a href='/2024/0918/c10846a353698/page.htm' target='_blank' title='计算机学院开展“开新课、新开课”试讲活动'>计算机学院开展“开新课、新开课”试讲活动</a><i>09-18</i></li>
                        
                      </ul>
                    </div> 
</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="col xs-12 sm-4 md-4 lg-4">
          <div class="section-column">
            <div class="in_title"><span>公示</span><a href="/10847/list.htm">更多>></a></div>
            <div class="list-ul list-ul-b" frag="窗口5" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'6','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/公示'}"><div id="wp_news_w5"> 

              <ul class="list-ul-li list-ul-li-b">
                
                <li><a href='/2024/0919/c10847a353834/page.htm' target='_blank' title='考察预告'>考察预告</a><i>09-19</i></li>
                
                <li><a href='/2024/0911/c10847a353216/page.htm' target='_blank' title='2024级人工智能创新班拟录取学生名单公示'>2024级人工智能创新班拟录取学生名单公示</a><i>09-11</i></li>
                
                <li><a href='/2024/0906/c10847a352631/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>09-06</i></li>
                
                <li><a href='/2024/0630/c10847a348271/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>06-30</i></li>
                
                <li><a href='/2024/0625/c10847a347930/page.htm' target='_blank' title='2023级“卓越工程师教育培养计划”专业 “卓越班”拟录取公示'>2023级“卓越工程师教育培养计划”专业 “...</a><i>06-25</i></li>
                
                <li><a href='/2024/0621/c10847a347604/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>06-21</i></li>
                
              </ul>
            </div> 
</div>
            <div class="zt">
              <div class="zt-title">
                <p>专</p>
                <p>题</p>
              </div>
              <div class="zt-con" frag="窗口18" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'1','c2':'图片路径','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'15','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/专题'}"><div id="wp_news_w18"> 

                
                <a href="http://47.100.138.90/SpaCCS2020/" target="_blank"><img src="/_upload/article/images/d7/71/813ceda34f738b57ac391a1734b7/6912482e-b750-4e38-bb29-dac7727f56d3.jpg"></img></a>
                
              </div> 
</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script type="text/javascript">
  $('.bxslider03').bxSlider({
    mode: 'horizontal',
    auto: true,
    autoControls: true,
    autoHover: true,
    pager: false,
    pause: 3000
  });
  </script>
  <div class="in-news">
    <div class="container container-pd">
      <div class="in_title"><span>通知公告</span><a href="http://cs.nuaa.edu.cn/tzgg/list.htm">更多>></a></div>
      <div class="row">
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>党委行政</span><a href="/dwxz/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口7" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/党委行政'}"><div id="wp_news_w7"> 

              <ul>
                
                <li><a href="/2024/0919/c10853a353834/page.htm" target="_blank"><i>></i>考察预告</a><span>09.19</span></li>
                
                <li><a href="/2024/0906/c10853a352631/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>09.06</span></li>
                
                <li><a href="/2024/0630/c10853a348271/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>06.30</span></li>
                
                <li><a href="/2024/0621/c10853a347604/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>06.21</span></li>
                
                <li><a href="/2024/0603/c10853a341740/page.htm" target="_blank"><i>></i>关于2022-2024年度学院先进基层党支部、优...</a><span>06.03</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>人事</span><a href="/10848/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口8" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/人事'}"><div id="wp_news_w8"> 

              <ul>
                
                <li><a href="/2023/1214/c10848a327388/page.htm" target="_blank"><i>></i>2023年度教职工考核工作安排</a><span>12.14</span></li>
                
                <li><a href="/2023/0302/c10848a303998/page.htm" target="_blank"><i>></i>计算机科学与技术学院/人工智能学院/软件学...</a><span>03.02</span></li>
                
                <li><a href="/2022/1210/c10848a300564/page.htm" target="_blank"><i>></i>2022年度教职工考核工作安排</a><span>12.10</span></li>
                
                <li><a href="/2022/0922/c10848a293100/page.htm" target="_blank"><i>></i>2022年拟申报初级专业技术职务人员公示</a><span>09.22</span></li>
                
                <li><a href="/2021/1214/c10848a272172/page.htm" target="_blank"><i>></i>【年终考核】2021年度计算机科学与技术学院...</a><span>12.14</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>学科科研</span><a href="/10849/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口9" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/学科科研'}"><div id="wp_news_w9"> 

              <ul>
                
                <li><a href="/2024/0420/c10849a336622/page.htm" target="_blank"><i>></i>关于组织模式分析与机器智能工业和信息化部...</a><span>04.20</span></li>
                
                <li><a href="/2024/0412/c10849a335982/page.htm" target="_blank"><i>></i>关于组织高安全系统的软件开发与验证技术工...</a><span>04.12</span></li>
                
                <li><a href="/2024/0410/c10849a335810/page.htm" target="_blank"><i>></i>关于组织脑机智能技术教育部重点实验室2024...</a><span>04.10</span></li>
                
                <li><a href="/2024/0102/c10849a328638/page.htm" target="_blank"><i>></i>关于开展第十八届中国青年科技奖候选人提名...</a><span>01.02</span></li>
                
                <li><a href="/2023/1227/c10849a328354/page.htm" target="_blank"><i>></i>关于开展第二十届中国青年女科学家奖和第九...</a><span>12.27</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>本科生培养</span><a href="/10850/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口10" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/本科生培养'}"><div id="wp_news_w10"> 

              <ul>
                
                <li><a href="/2024/1012/c10850a355614/page.htm" target="_blank"><i>></i>【创新班专业分流】关于公布计算机科学与技...</a><span>10.12</span></li>
                
                <li><a href="/2024/1009/c10850a355263/page.htm" target="_blank"><i>></i>【创新班考核】2023-2024 学年 2021、2022...</a><span>10.09</span></li>
                
                <li><a href="/2024/0925/c10850a354395/page.htm" target="_blank"><i>></i>【创新班考核】关于开展2021级、2022级、20...</a><span>09.25</span></li>
                
                <li><a href="/2024/0916/c10850a353544/page.htm" target="_blank"><i>></i>【推免工作】2025届本科毕业生拟推荐免试攻...</a><span>09.16</span></li>
                
                <li><a href="/2024/0912/c10850a353274/page.htm" target="_blank"><i>></i>【推免工作】2025届本科毕业生推荐免试攻读...</a><span>09.12</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>研究生培养</span><a href="/10851/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口11" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/研究生培养'}"><div id="wp_news_w11"> 

              <ul>
                
                <li><a href="/2024/1015/c10851a355802/page.htm" target="_blank"><i>></i>【博士答辩】董晨刚 博士答辩公告</a><span>10.15</span></li>
                
                <li><a href="/2024/1008/c10851a355112/page.htm" target="_blank"><i>></i>【博士答辩】石优 博士答辩公告</a><span>10.08</span></li>
                
                <li><a href="/2024/0926/c10851a354484/page.htm" target="_blank"><i>></i>【博士答辩】耿琳 博士答辩公告</a><span>09.26</span></li>
                
                <li><a href="/2024/0926/c10851a354483/page.htm" target="_blank"><i>></i>【博士答辩】吴平凡 博士答辩公告</a><span>09.26</span></li>
                
                <li><a href="/2024/0926/c10851a354482/page.htm" target="_blank"><i>></i>【博士答辩】田苗 博士答辩公告</a><span>09.26</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>学生工作</span><a href="/10852/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口12" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/学生工作'}"><div id="wp_news_w12"> 

              <ul>
                
                <li><a href="/2024/0712/c10852a350851/page.htm" target="_blank"><i>></i>我院举办“生涯辅导谈话主题工作坊”暨班主...</a><span>07.12</span></li>
                
                <li><a href="/2024/0613/c10852a342859/page.htm" target="_blank"><i>></i>计算机学院开展“浓情端午‘粽’享美好”端...</a><span>06.13</span></li>
                
                <li><a href="/2024/0529/c10852a341300/page.htm" target="_blank"><i>></i>计算机学院举办“导学引领，乐享运动”2024...</a><span>05.29</span></li>
                
                <li><a href="/2024/0528/c10852a341072/page.htm" target="_blank"><i>></i>计算机学院举办“生涯筑梦 计遇未来” 职业...</a><span>05.28</span></li>
                
                <li><a href="/2024/0424/c10852a337149/page.htm" target="_blank"><i>></i>翼下山河，守护疆土——飞行员的航空梦想与爱...</a><span>04.24</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script src="/_upload/tpl/04/02/1026/template1026/js/owl.carousel.min.js"></script>
  <script type="text/javascript">
  jQuery("#owl-demo1D").owlCarousel({
    loop: false,
    margin: 15,
    autoWidth: true,
    items: 3,
    dots: false,
    mouseDrag: false
  });
  $(".owl-item").on('click', '.item .news_title_item', function() {
    $('.news_title_item').removeClass('active');
    $(this).addClass('active');
    var index = $(this).parents(".owl-item").index();
    $('.in-news-con .row').eq(index).stop().slideDown().siblings().stop().slideUp();
  });
  </script>
  <div class="in-info">
    <div class="container container-pd">
      <div class="in_title in_title_c in-info-span"><span class="active">学术信息</span></div>
      <div class="in-info-con">
        <div class="row" style="display: block;" frag="窗口13" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'6','c2':'标题,扩展字段11,扩展字段12,扩展字段13,扩展字段14,扩展字段15','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'30','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'0','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/学术信息'}"><div id="wp_news_w13"> 

          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-10-18</div>
              </div>
              <div class="title"> <a href='/2024/1017/c11617a355952/page.htm' target='_blank' title='【学术会议】钟山论坛—网络与信息安全前沿技术研讨会'>【学术会议】钟山论坛—网络与信息安全前沿...</a>
                <p class="p1">题目：钟山论坛—网络与信息安全前沿技术研讨会</p>
                <p class="p2">报告人：刘哲理、纪守领、沈剑、马小博</p>
                <p class="p3">时间： <span class="t-time">2024-10-18</span> 下午14:50-17:10</p>
                <p class="p4">地点：计算机学院楼515会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-10-28</div>
              </div>
              <div class="title"> <a href='/2024/1010/c11617a355363/page.htm' target='_blank' title='【石榴大讲堂】深度学习系统的可信性保障'>【石榴大讲堂】深度学习系统的可信性保障</a>
                <p class="p1">题目：深度学习系统的可信性保障</p>
                <p class="p2">报告人：孙猛</p>
                <p class="p3">时间： <span class="t-time">2024-10-28</span> 上午10点</p>
                <p class="p4">地点：计算机学院507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-09-27</div>
              </div>
              <div class="title"> <a href='/2024/0923/c11617a354088/page.htm' target='_blank' title='【石榴大讲堂】面向机器学习的分布式优化算法设计与分析'>【石榴大讲堂】面向机器学习的分布式优化算...</a>
                <p class="p1">题目：面向机器学习的分布式优化算法设计与分析</p>
                <p class="p2">报告人：杨绍富</p>
                <p class="p3">时间： <span class="t-time">2024-09-27</span>  上午11:00 - 12:00</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-09-27</div>
              </div>
              <div class="title"> <a href='/2024/0923/c11617a354087/page.htm' target='_blank' title='【石榴大讲堂】通信资源受限下完全分布式协同控制与博弈'>【石榴大讲堂】通信资源受限下完全分布式协...</a>
                <p class="p1">题目：通信资源受限下完全分布式协同控制与博弈</p>
                <p class="p2">报告人：许文盈</p>
                <p class="p3">时间： <span class="t-time">2024-09-27</span> 上午10:00-11:00</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-09-26</div>
              </div>
              <div class="title"> <a href='/2024/0920/c11617a353885/page.htm' target='_blank' title='【石榴大讲堂】Quantum Software Engineering- A New Genre of Computing'>【石榴大讲堂】Quantum Software Engineeri...</a>
                <p class="p1">题目：Quantum Software Engineering- A New Genre of Computing</p>
                <p class="p2">报告人：Arif Ali Khan</p>
                <p class="p3">时间： <span class="t-time">2024-09-26</span> 下午3:00-4:00</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-09-23</div>
              </div>
              <div class="title"> <a href='/2024/0919/c11617a353819/page.htm' target='_blank' title='【石榴大讲堂】大模型中数据合成技术的应用'>【石榴大讲堂】大模型中数据合成技术的应用</a>
                <p class="p1">题目：大模型中数据合成技术的应用</p>
                <p class="p2">报告人：王科</p>
                <p class="p3">时间： <span class="t-time">2024-09-23</span> 下午3:00-4::00</p>
                <p class="p4">地点：计算机学院楼511会议室</p>
              </div>
            </div>
          </div>
          
          <a class="in-news-more" href="/10854/list.htm">更多>></a>
        </div> 
</div>
      </div>
    </div>
  </div>
  <script type="text/javascript">
  $(function() {
    // 文章日历
    $(".rili").sudyPubdate({
      target: ".date", // 日期元素
      lang: "en", //    月份语言，支持"en"英文，"cn"中文, "num"数字，默认为数字
      separator: "-", // 文章日期之间分隔符的，默认是后台输出的"-"号
      format: "年月日", // 文章日期格式，默认为年月日
      tpl: '<h3>%m%</h3><p>%d%</span></p>'
      // 自定义输出格式  %Y%代表年，%m%代表月, %d%代表日, %w%代表星期, %M%代表翻译后月份，%W%代表翻译后星期
    });
    // 文章日历
    $(".p3").sudyPubdate({
      target: ".t-time", // 日期元素
      lang: "", //  月份语言，支持"en"英文，"cn"中文, "num"数字，默认为数字
      separator: "-", // 文章日期之间分隔符的，默认是后台输出的"-"号
      format: "年月日", // 文章日期格式，默认为年月日
      tpl: '%m%月%d%日'
      // 自定义输出格式  %Y%代表年，%m%代表月, %d%代表日, %w%代表星期, %M%代表翻译后月份，%W%代表翻译后星期
    });
  });
  </script>
  <script type="text/javascript">
  $('.in-info-span span').on('click', function() {
    var index = $(this).index();
    $(this).addClass('active').siblings().removeClass('active');
    $('.in-info-con .row').eq(index).show().siblings().hide();
  });
  </script>
  <div class="lnk">
    <div class="container">
      <div class="row">
        <div class="col xs-12 sm-12 md-5 lg-5">
          <div class="lnk-l">
            <div class="row">
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="http://newsweb.nuaa.edu.cn/"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a1.png"></span>
                    <p>新闻网</p>
                  </a> </div>
              </div>
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="http://i.nuaa.edu.cn/"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a2.png"></span>
                    <p>智慧校园</p>
                  </a> </div>
              </div>
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="http://lib.nuaa.edu.cn/"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a3.png"></span>
                    <p>图书馆</p>
                  </a> </div>
              </div>
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="https://mail.nuaa.edu.cn/coremail/index.jsp"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a4.png"></span>
                    <p>电邮系统</p>
                  </a> </div>
              </div>
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="http://hqjt.nuaa.edu.cn/bcsk/list.htm"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a5.png"></span>
                    <p>班车时刻</p>
                  </a> </div>
              </div>
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="https://oas.nuaa.edu.cn"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a6.png"></span>
                    <p>OA办公网</p>
                  </a> </div>
              </div>
            </div>
          </div>
        </div>
        <div class="col xs-12 sm-12 md-7 lg-7">
          <div class="lnk-r">
            <div class="row">
              <div class="col xs-6 sm-3 md-3 lg-3">
                <div class="lnk-r-item"> <a href="/2017/1115/c11649a177680/page.htm" target="_blank"><img src="/_upload/tpl/04/02/1026/template1026/images/b1.jpg"></a> </div>
              </div>
              <div class="col xs-6 sm-3 md-3 lg-3">
                <div class="lnk-r-item"> <a href="http://keyselab.nuaa.edu.cn/"><img src="/_upload/tpl/04/02/1026/template1026/images/b2.jpg"></a> </div>
              </div>
              <div class="col xs-6 sm-3 md-3 lg-3">
                <div class="lnk-r-item"> <a href="http://pami.nuaa.edu.cn/"><img src="/_upload/tpl/04/02/1026/template1026/images/b3.jpg"></a> </div>
              </div>
              <div class="col xs-6 sm-3 md-3 lg-3">
                <div class="lnk-r-item"> <a href="javascript:void(0)"><img src="/_upload/tpl/04/02/1026/template1026/images/b4.jpg"></a> </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="foot">
    <div class="container container-pd">
      <div class="fo-con clearfix">
        <div class="row">
          <div class="col xs-12 sm-7 md-8 lg-9">
            <div class="fo-l">
              <p>地址：江苏省南京市江宁区将军大道29号</p>
              <p>邮政编码: 211106 </p>
              <p>版权所有：南京航空航天大学计算机科学与技术学院/人工智能学院 ALL RIGHTS RESERVED 苏ICP备05070685号 <a href="http://site.nuaa.edu.cn/" target="_blank">后台管理</a> <a href="mailto:njliujr@nuaa.edu.cn" target="_blank"><img src="/_upload/tpl/04/02/1026/template1026/images/email.png"> 书记信箱</a> <a href="mailto:cb_china@nuaa.edu.cn" target="_blank"><img src="/_upload/tpl/04/02/1026/template1026/images/email.png"> 院长信箱</a></p>
            </div>
          </div>
          <div class="col xs-12 sm-5 md-4 lg-3">
            <div class="fo-r">
              <h3>友情链接</h3>
              <div class="fo-r-con clearfix">
                <div class="select fl" frag="窗口15" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'30','c2':'标题','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'-1','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/友情链接'}"><div id="wp_news_w15"> 

                  <p>校外导航链接<i></i></p>
                  <ul>
                    
                    <li><a href='javascript:void(0)' target='_blank' title='友情链接'>友情链接</a></li>
                    
                  </ul>
                </div> 
</div>
                <div class="fo-lnk clearfix fl"> <a class="wx" href="javascript:;">
                    <div class="ewm"><img src="/_upload/tpl/04/02/1026/template1026/images/ewm.png" /></div>
                  </a> <a class="wb" href="javascript:void(0)"></a> </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script>
  $(function() {
    $(".select p").on("click", function(e) {
      var objDiv = $(this).next('ul');
      if (objDiv.css('display') == 'none') {
        objDiv.css('display', 'block');
        objDiv.parent().siblings().find('ul').css('display', 'none');
        event.stopPropagation();
      } else {
        objDiv.css('display', 'none');
      }
      $(document).one("click", function() {
        objDiv.hide();
      });
      e.stopPropagation();
    });
    $(".select p").next('ul').on("click", function(e) {
      e.stopPropagation();
    });
    $('.wx').on('click', function() {
      $('.ewm').toggle();
    });
  });
  </script>
</body>

</html>
 <img src="/_visitcount?siteId=68&type=1&columnId=1952" style="display:none" width="0" height="0"/>"
	fmt.Println(s)
}
