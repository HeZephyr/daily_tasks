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

<meta content="4oMaBbpgiODhKCRYDXCVO97tHvugbVaY" r="m"><!--[if lt IE 9]><script r='m'>document.createElement("section")</script><![endif]--><script type="text/javascript" r='m'>$_ts=window['$_ts'];if(!$_ts)$_ts={};$_ts.nsd=15896;$_ts.cd="qxzDrrAlDPLllqVmtGQArAAjqfqmqu7qtGqLqpq5laV7cqAXiaqjqnLqlOg4qcLDtGQAHO7qtGVjqfE5iaqjqcLrlkg4qcgPcs7ctG3XEfVHhq9lEcLclqVkiAqjqcLrlkg4qcLotGAAHO9qtGEjqSq5tGq7cGA2tGE1rAW4qcLqtGlAHPLllqVmiAqjqPLklkgjrfqmqSAltGVrhcFPqs3aqGVnJFEOtLXKjS7K5Q272l4z9VfnAtVar1__DuNtpbGG95sjVFgqrqEorqcJkGRI81VC8i0aWoW7YCYtZ022p9fCIVONsbymEkTIKB9fUPzOw6A7TUqBQ6lLRUo1hvJfQomz37Q7hPgGt6ruTCRwK1yNQK6ihCySMb7LMewBMcf7QbfzObTTFDg.Jc6WwbSLUn7bAN7u3CzGKmzy4OREHl3Ns6o1psVZp1V5K7SCMDGXFUmzuPeXQKy.tDdUMCaLMoxLM3yzwKfXtuAzuoR7FmZ.x9UkWkxXVsy.iirEWVZ0YOSEnKTjYoW2H2.WwbSLhb2SMQgBMvmXFnfBTCe7tDSnFDCiMUYLMP70hQSCMDfItnRA4KYEMUmpMC.kJsGjwmTUKWRbwmROxsyFuoR7FcyNQK6ihCySMb7LMewBMcf7QbfzObTTFDg.Jc6WwbSLUn7bV7Sx8vw0F6xM2D2TYuYOY9kJibTMp1V5K7SCMDGXFUmzuPeXQKy.tDdUMCaLMoxLM3yzwKfXtuAzuoR7FmZ.x2BFR9xz8ue7FQrxhPNuwKelb2x_JlV2H2.WwbSLhb2SMQgBMvmXFnfBTCe7tDSnFDCiMUYLMP70hQSCMDfItnRGnuN9Q02iFmIJsO2fMmVTKWzqKsTmxsyFuoR7FcyNQK6ihCySMb7LMewBMcf7QbfzObTTFDg.Jc6WwbSLUn7b3ZyA8sa7R20022RpWsNGFutUsDJzAnV5K7SCMDGXFUmzuPeXQKy.tDdUMCaLMoxLM3yzwKfXtuAzuoR7FmZ.xbjKMYeGQUluWNN0VupgU9zK4l0yK0W2H2.WwbSLhb2SMQgBMvmXFnfBTCe7tDSnFDCiMUYLMP70hQSCMDfItnRB4DNKMCTG3KHtYbeeQbyK3zfBQKzkxsyFuoR7FcyNQK6ihCySMb7LMewBMcf7QbfzObTTFDg.Jc6WwbSLUn7bMdm9MVJqU0y3S92BpuJHMVdVAOYO3nV5K7SCMDGXFUmzuPeXQKy.tDdUMCaLMoxLM3yzwKfXtuAzuoR7Fme_tm.yMUYLMPzjwWyBhc3nx1GObOWPtcrRIUBptCxNIbrLtESzwKYOtKm2TnTOFPq.FvsiMna0hb2SMQy8hme7QbfzObTTFDg.xODqhnEZWPVLEda.wD2jMDxu6vwNRKS4wCOrM6S63Uz63WSGFomOFKJe062ZwoNdQo61w6S0QDyjwwyuFoTSMb2e0bzy3UN68D4WRvS.Foz.RwmdMCYTw6weuopa8DSnRKkMMoYaFvzjQ7wdMbSBIP0BuoRfR1SP3o1W3Dgbhb2SMQgBWcf7QbfzdneMFvw.Fc6WwbSLhPV0EggfikAGtnRIPbrf3KY.8oIi3bxgwKxz8Qzzt1N7QbYGODR9Q1S9FPiiMUYLMP70hQSCMDfItTeBTCe7tDSnFDCiEsqbhPVZW3qBE2LjQC2NCDmZQoR98oskMvSnF6YawWJBQCTB86xXSvr2wvpPwv6VQKYjRKyjRJy0Fop5QDTBSvwTFbSPFo6YQnyX3Uzu8Q2_8or5QoyvTUegJvydR66F3DRPIveOIwNXRCG.tKN0aCEX3KRntbhWE1SjwCzLhREBMvmXFT9z5KTTFDg.FvsiMnabWPVLEyLuEPGGK1SOCDz.3oy7QovrtPyjwCR2hZwvw1NfF13zuoR7Fcg6tDdUMCSRh2SjwWyBhDN6FCGzGkWPtcqdJciiEYg7RozP3Wmbt1N7QbYGODR9Q1S9FPiiMUYLMP70hQSCMDfItTeBTCe7tDSnFDCiEsqbhPVZW3qBE2Lj3bzz6bzfIURPw6KhhKyS3C3jRWRChbw7x1fBTCe7tkE.FvsiMT0LKK2SMQgBMvmXFnGO9Pl7xOL6xPCqU1f5RowNRJmdRKTv8C2gSbx0RUyj3K6k3USbFCfGMzmjMKmCwUTeSDRfRKpOw66tQUSOMK2OQ7Njt1N7QbYGODR9Q1S9FPiiMUYLMP70hQSCMDfItTeBTCe7tDSnFDCiEsqbhPVZW3qBE2LjIKrP0UeuMbmz8oUWwbm2IbTjwWJPIoybIC3LODTTRKWN3KUUhKJjEczjwWyBhkAXFUmzu2g7UCSnFDCiMUYLMP7bW3qBEOanx1GOdczfMUr9wbklwnT2QUpfMWrB8UfO3C2L0UrNFCS08DkrMvS23KJNIQJPMn2vRvQLuoRfR1SP3o1W3Dgbhb2SMQgBWcf7QbfzdneMFvw.Fc6WwbSLhPV0EggfikAGtnRI0UwThDRPFvU8tPyjwCR2hZwvw1NfF13zuoR7Fcg6tDdUMCSRh2SjwWyBhDN6FCGzGkWPtcqdJciiEYg786egwwVjhbN6RbVBSCwTtbmNxP6WwbSLhOQLMewBMm9XUDN0ubZ7Fvw.FcCqWnELEkZ0EggfUPSjFomnS6e7Ioy0IUhYRbxgw6Y0WE7zMvmORPN9Sv3XRCa2tDdUMCaLWPzjwWyBU1fRFUmzuPeXQKy.tci1E1abiOQbh3rwtow5wDSG0D3NtbSnRKcWRbpShDpjEgyzwKfXtuAzuoR7FmZ.UCdUMCaLMoxLM3gfWc3XxsaSGcZPKPNGIv6m8bRgRUegRzwPMbf.8CJuaCwTRURP8DMq3orgRoygRew0wUf9Q6ST6bYfR6w.QoPAICR23UxLwwzdRbYfQbfTCny2FvyP3DIHQ6Sf3DyLMwyv8vS_3DweSvzaIoyOQDhr3vSGF6p6tESzwKYOtKm2TnTOFPq.FvsiMna0hb2SMQy8hme7QbfzObTTFDg.xODqhnEZWPVLEda.8brfwof.uoRfRUyuFvsk3vS6FDpXIwfv8D3.tKN0aCEX3KRntbhWE1SjwCzLhREBMvmXFT9z5KTTFDg.FvsiMnabWPVLEyLuEPGGK1S.aKeNRvYPMUvhMUSX3Kzz3eyN3Cy_8CyPn6Jfw6m4MbhqQoNaIbTvIQzaRUf.wUfLuUwuR6yXF6KAIC2.wORgFWT6IDT7wofLu6rZMKybRv6h3begFvyLwzxaIDSCFDYeSU2PIoxuh1dWwbm2hDxGwESGMP3XFUmzuPZatDSnFD6dhTzjwCzLhQSCMDGXxsAOOPlgJcq.x25t8UJ5M6z5QzmbQvx.RDmeCoJZIb2dQo63F6fOIvT7R7Nd8bJ5IKxOu6e_3oy5QK.owCmjRUe6Iwz6RKYB86Tn0KmZIK2utvBDI6feFvev8wy48KSjQ6fy0Kpy8oN03nPWMUY23n2fRwQz3CLGtCN0ubZ7JcyNQK6iUPSIMoxLM3yzwKfXtn3SGcZPHkE2tcHZt6r786pz3zmdRKfvFCYOPnTXQKJbtbsDwPyvMcVLMewBMcGntCN0ubewtm2NQK6ihCySMb7LEyEfhc3_Jn3zGm0LFKmjRbu8MbYL3CYe8E7zMvmORPN9Sv3XRCa2tDdUMCaLWPzjwWyBU1fRFUmzuPeXQKy.tci1E1abiOQbh3rwtDSvwUf40bgNtbSnRKcWRbpShDpjEgyzwKfXtuAzuoR7FmZ.UCdUMCaLMoxLM3gfWc3XxsaSGcZPtPGNFvsk3PyfRvEj37afhDN6FCGz9PeXQKy.K16yMUYLMPzjwWyBhc3nx1GObOWPtcrRhoh13bJfhCxZQQy7QbYCMbzB6vRfQvmaRUHHICmfQDffMZYjMblBFCYO6bNWRKqjIl6OAmYrpvz2RJr7RUpS86xL0owTRKy0MKnhhKyS3C3jRWRChbw7x1fBTCe7tkE.FvsiMT0LKK2SMQgBMvmXFnGO9Pl7xOL6xPCqU1fPFC2GIQyaQofB3KTBSPLXFvwbR1doR6Vj3K9bhQSCMDGXJnfBTCe7K1yQFvsiMnSjwCzLh3quEPGGHuAOOPm8hDN6RCHiFUS631gjMewP31Nv36lBaK0PtDSnFDCiWnSjwCzLUEy3MvmXFnfBTCe7tcq6xPCqiuqbhPYFQQJ63Cm7QCmuPPTOFKwP8cPWMUY23n2fRwQz3CLGtCN0ubZ7JcyNQK6iUPSIMoxLM3yzwKfXtn3SGcZPHkE2tcHZtCmLQUJSIwNaRCG.tKyPODTTRKWN3KUUhKJjEczjwWyBhkAXFUmzu2g7UCSnFDCiMUYLMP7bW3qBEOanx1GOdczXMDNC8Ddt3DYgMDxaIQebwofB3bmTnCTv8DyXRvUoQopgMbybQZT7MUfXMbRuS6Wn8DyXRvHoQ6SLFCYO3JRbIDf.RUfjOc2ZMvNGFcdp3D2fMDpfIQf4R6fLFDRenDpu8DfOJsD1tPyjwCR2hZwvw1NfF13zuoR7Fcg6tDdUMCSRh2SjwWyBhDN6FCGzGkWPtcqdJciiEYg73UJPIQmahUSbwofP01y7Q6x08Dh1ICJNRCxeQwzd3Ce_RKR76brv8DmPFKCHwUxg3KxzRWT2RbTf3vfPS6eORCJbF6H83Dx6IbY.M72GQbz5RvwT0DazFCJ08Dks8UJZ8vz2QWfGIDYXQb2Bn6pgh1SNQKkkhKYGwn2vMgqBMvmXFnGSObTTFDyFKYnLqGyQFueP3zrrsOZnJumlTUATWO20V9OlWDwq1uSXpdxNJuSV1keG_lfiYoYj1s.vVlYbWbTK8ZeUW9Wy10TKSKz1VYxmWmkNim2T36JmMZ2rQ6JGFke14OYRWoffATsoR0QjWuge37An3PNLwkeXC2N1pCwYJKBERlZjWuge37An3PNLwkeHb2xVFCJjRTsqY0ljWuge37An3PNLwkeCZvRhs2R6R2kh3O9jFoAdR59SwOrgIbJPT9mFHCJlUkIkWu2OW2rDFx2upvwEAla0TbzLWOAZM9M6popDpvJE37wphOV4Hbwv9ClXMvlZFuKsROxqV0m6s7w7hOV4Hbwv9ClXMvlZtOcMibJPWCVjFelkqqAkJaAq9aWUMoRuwkBQh1LjMox23ESbRol7RDL4bOWqqOW5HKhlWGqoWulCqFGnWkWCWk0S2OQarqVdJuPAJkqaqf8JFe5s7xq5vh_Fc0JbEQzJrrDpfqg.734MmtifqaADrAWmLPHsuYwq0iMozTQX0QZ01ZgjN.rSQ2.u33oZh6P6R9kHUAqiWs3aWRE4isE4JaAq9aWtJsV0WO1UJsAnJAQBMyzA1sm2QvQuS2pmAVTa3VP1A2rZUUeT1iyN3CJGpKyR4lwWQ0w3IVHtiGqlqGQlqxEkqsWk5kJ4ObxfRnybFUii3vJNhbRnFEyf3KZXRKYTObmzwcy2MK5i3cSv3CVL3zmahDwSQ1fPuKg7R6x6tDUkwnSGMb7LRQTGhDrzQnf2uKrgtDwd316oMoALRDz2hQpB3nf9FDLzSDS0tDpLFvKiRDpbhbrzRgy98Pf2MC7znbeatDN6wc6twCWLFCp.hQzzF1f.FKEznCTutDz0Fn6EQcSX36VLFex9hDyn3bWznorBFPyzRUIkhCSz3czLM7VBMKYntC2LSPezwk3.FUokhC0a3nzNQRlBMUp2tC2jTcezFC3.FK4VhCTzMnzNQwEBMUw.tC2SSCA7FCrztD4r8nS.RKqLM7yfhDzX3PfXTbV7wo7.woOrh6r2Q1z0R7ZBQopBt6pS01eawo3.wDUlh6r0wvZLQQS9hop.RcfSubL7womOtokhR1STwvZLQZWBQb9XwKY7OvmzQbq.wbsmh6xfQ1zCwQABQbYat6wXOvrPFcy03vbiQow2hvpuRgyawbWXwoTnOvwfwPyuRUXiwC2PhvJzRZABwDzawKmBTPe0Io3.QoUlh6R6RPznQzlBwvwTt6xySPe6FCpftoMJRK3ChvrNwgy0MKVXQDY6OvpawnydRbni86rNhveLQLy4Mow_t6SnuPegFKg.IP63FvqL36wBhwxBFUwCtCTSane7RKqorqDbUcpMRGQqUL3vpGADKcrPuGWlKnRb1qDGWkqlqkq6qFWnqqVuraArLG";if($_ts.lcd)$_ts.lcd();</script><script type="text/javascript" charset="utf-8" src="/LIWghRUPB4QK/oxYRCeR1jjgO.199cf1b.js" r='m'></script><script language="javascript" src="/_js/sudy-jquery-autoload.js" jquery-src="/_js/jquery-1.9.1.min.js" sudy-wp-context="" sudy-wp-siteId="68"></script>
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
      <div class="t_se clearfix" frag="窗口0" portletmode="search">            <form method="post" action="/_web/_search/api/search/new.rst?locale=zh_CN&request_locale=zh_CN&_p=YXM9NjgmdD0xMDI2JmQ9NDE2MCZwPTEmbT1TTiY_" target="_blank">
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
            <div class="top_nav hidden-xs" frag="窗口1" portletmode="simpleSudyNavi" configs="{'c1':'1','c7':'2','c4':'_self','c3':'30','c6':'0','c8':'2','c9':'0','c2':'1','c5':'2'}" contents="{'c2':'0', 'c1':'/学院概况,/师资队伍,/科学研究,/教学工作,/学生工作,/党群工作,/资产管理'}">
              
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
                    
                    <li><a href="/19939/list.htm" target="_self">科研概况</a></li>
                    
                    <li><a href="/19940/list.htm" target="_self">科研项目</a></li>
                    
                    <li><a href="/19941/list.htm" target="_self">科研经费</a></li>
                    
                    <li><a href="/19942/list.htm" target="_self">科研成果</a></li>
                    
                    <li><a href="/19943/list.htm" target="_self">学术交流</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1958/list.htm" target="_self">教学工作</a>
                  
                  <ul>
                    
                    <li><a href="/zyjs/list.htm" target="_self">专业介绍</a></li>
                    
                    <li><a href="/jxzd/list.htm" target="_self">教学制度</a></li>
                    
                    <li><a href="http://cs.nuaa.edu.cn/10850/list.htm" target="_self">本科生培养</a></li>
                    
                    <li><a href="http://cs.nuaa.edu.cn/10851/list.htm" target="_self">研究生培养</a></li>
                    
                    <li><a href="/1981/list.htm" target="_self">资料下载</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1959/list.htm" target="_self">学生工作</a>
                  
                  <ul>
                    
                    <li><a href="/19944/list.htm" target="_self">学生活动</a></li>
                    
                    <li><a href="/19945/list.htm" target="_self">党团建设</a></li>
                    
                    <li><a href="/19946/list.htm" target="_self">访企拓岗</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1960/list.htm" target="_self">党群工作</a>
                  
                  <ul>
                    
                    <li><a href="/19947/list.htm" target="_self">学院党委委员会</a></li>
                    
                    <li><a href="/19948/list.htm" target="_self">学院工会委员会</a></li>
                    
                    <li><a href="/19949/list.htm" target="_self">通知公告</a></li>
                    
                    <li><a href="/19950/list.htm" target="_self">工作动态</a></li>
                    
                    <li><a href="/19951/list.htm" target="_self">学习园地</a></li>
                    
                    <li><a href="/19952/list.htm" target="_self">工作文件</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/19916/list.htm" target="_self">资产管理</a>
                  
                </li>
                
              </ul>
              
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="phone_menu_t visible-xs">
    <div class="phone_menu" style="display: none" frag="窗口01" portletmode="simpleSudyNavi" configs="{'c1':'1','c7':'2','c4':'_self','c3':'30','c6':'0','c8':'2','c9':'0','c2':'1','c5':'2'}" contents="{'c2':'0', 'c1':'/学院概况,/师资队伍,/科学研究,/教学工作,/学生工作,/党群工作,/资产管理'}">
      
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
            
            <li><a href="/19939/list.htm" target="_self">科研概况</a></li>
            
            <li><a href="/19940/list.htm" target="_self">科研项目</a></li>
            
            <li><a href="/19941/list.htm" target="_self">科研经费</a></li>
            
            <li><a href="/19942/list.htm" target="_self">科研成果</a></li>
            
            <li><a href="/19943/list.htm" target="_self">学术交流</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1958/list.htm" target="_self">教学工作</a>
          
          <ul>
            
            <li><a href="/zyjs/list.htm" target="_self">专业介绍</a></li>
            
            <li><a href="/jxzd/list.htm" target="_self">教学制度</a></li>
            
            <li><a href="http://cs.nuaa.edu.cn/10850/list.htm" target="_self">本科生培养</a></li>
            
            <li><a href="http://cs.nuaa.edu.cn/10851/list.htm" target="_self">研究生培养</a></li>
            
            <li><a href="/1981/list.htm" target="_self">资料下载</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1959/list.htm" target="_self">学生工作</a>
          
          <ul>
            
            <li><a href="/19944/list.htm" target="_self">学生活动</a></li>
            
            <li><a href="/19945/list.htm" target="_self">党团建设</a></li>
            
            <li><a href="/19946/list.htm" target="_self">访企拓岗</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1960/list.htm" target="_self">党群工作</a>
          
          <ul>
            
            <li><a href="/19947/list.htm" target="_self">学院党委委员会</a></li>
            
            <li><a href="/19948/list.htm" target="_self">学院工会委员会</a></li>
            
            <li><a href="/19949/list.htm" target="_self">通知公告</a></li>
            
            <li><a href="/19950/list.htm" target="_self">工作动态</a></li>
            
            <li><a href="/19951/list.htm" target="_self">学习园地</a></li>
            
            <li><a href="/19952/list.htm" target="_self">工作文件</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/19916/list.htm" target="_self">资产管理</a>
          
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
  <div class="banner hidden-xs" frag="窗口2" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'6','c2':'图片路径','c27':'320','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'320','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'15','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'1', 'c1':'/PC首页幻灯'}"><div id="wp_news_w2"> 

    <ul class="bxslider">
      
      <li><img src="/_upload/article/images/71/0c/7c6f808e4cafa24a7054c4ad5135/3b79b3d9-e96c-4620-8018-ba69fa8444e4.png" /></li>
      
      <li><img src="/_upload/article/images/8a/a1/5dd77500459c86cb1c03521e709c/aa88887a-36fc-47e9-b371-18fee32f28c6.jpg" /></li>
      
      <li><img src="/_upload/article/images/19/aa/361b28a842a6ae3e34f2cba9ea67/ed252aef-483d-4998-a88f-88624607aeaa.jpg" /></li>
      
      <li><img src="/_upload/article/images/01/04/523d814246dda4c114c9425dce30/fd8e6513-4264-4a38-9b99-567e8a804b4c.png" /></li>
      
      <li><img src="/_upload/article/images/2c/68/7bd85cd148e287ce7296d350a07f/dd5534aa-c058-43bc-86f3-b03d80406e03.png" /></li>
      
      <li><img src="/_upload/article/images/de/1f/d41dead342b98bcd2fc4628e1067/681fed19-48c5-4eed-9828-01b0a69e130d.jpg" /></li>
      
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
                          <div style="background-image:url(/_upload/article/images/1b/83/cb663a374aa88d41150ab65b1499/3cd7d56f-a92f-4520-9d8d-33e0d1e05d56.jpg);"></div>
                          <p><a href='/2025/0107/c11624a362527/page.htm' target='_blank' title='教务处党支部与学院2021级本科生党支部开展联学共建党日活动'>教务处党支部与学院2021级本科生党支部开展联学共建党日活动</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/70/0a/4bfdf7584c799806664771d27f32/bcdfc403-3536-42ba-843f-743a7e75020e.jpg);"></div>
                          <p><a href='/2024/1223/c11624a361472/page.htm' target='_blank' title='学院举行“厚植爱国情怀，砥砺强国之志”主题升旗仪式'>学院举行“厚植爱国情怀，砥砺强国之志”主题升旗仪式</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/50/52/6eb87545481996fdc6764e85b8cd/de564509-d905-481a-8c00-9b7e0d005403.jpg);"></div>
                          <p><a href='/2024/1223/c11624a361470/page.htm' target='_blank' title='六WEI空间 | 包饺话冬至，合乐共此“食”'>六WEI空间 | 包饺话冬至，合乐共此“食”</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/91/d5/82fd20364592acd6b52b8b6d2c71/249a39c5-249c-4aa6-961b-39d6b011f760.jpg);"></div>
                          <p><a href='/2024/1209/c11624a360348/page.htm' target='_blank' title='学院举办第十七期“山湖号”红色专列活动'>学院举办第十七期“山湖号”红色专列活动</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/4c/f8/ea3064ee4682b7fbb79468d513de/6e294d5b-0de4-4e43-bb65-1ebb5ec4b16d.jpg);"></div>
                          <p><a href='/2024/1209/c11624a360335/page.htm' target='_blank' title='六WEI空间 | 榜样·对话，共享青春高光时刻'>六WEI空间 | 榜样·对话，共享青春高光时刻</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/39/fc/dd53981f470e9cd9509d1993a5d0/dfe6f739-ba2d-4063-a913-ea1d546b8aa6.jpg);"></div>
                          <p><a href='/2024/1209/c11624a360329/page.htm' target='_blank' title='六WEI空间 | 强国有我，青春有为'>六WEI空间 | 强国有我，青春有为</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/51/d8/fc9efbf44b8a855a24afdde4a20d/174a8049-152e-4ce6-8edd-c281955e640f.jpg);"></div>
                          <p><a href='/2024/1209/c11624a360327/page.htm' target='_blank' title='学生处党支部与学院本科生党支部开展“三个一”联系活动'>学生处党支部与学院本科生党支部开展“三个一”联系活动</a></p>
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
                        
                        <li><a href='/2025/0110/c10846a363037/page.htm' target='_blank' title='集美大学计算机工程学院来我院调研交流'>集美大学计算机工程学院来我院调研交流</a><i>01-10</i></li>
                        
                        <li><a href='/2025/0107/c10846a362527/page.htm' target='_blank' title='教务处党支部与学院2021级本科生党支部开展联学共建党日活动'>教务处党支部与学院2021级本科生党支部开展...</a><i>01-07</i></li>
                        
                        <li><a href='/2025/0102/c10846a362016/page.htm' target='_blank' title='我校成功举办第十五届本科生学术论坛'>我校成功举办第十五届本科生学术论坛</a><i>01-02</i></li>
                        
                        <li><a href='/2024/1225/c10846a361596/page.htm' target='_blank' title='学院工作案例入选2024年度“清廉南航”十佳案例'>学院工作案例入选2024年度“清廉南航”十佳...</a><i>12-25</i></li>
                        
                        <li><a href='/2024/1223/c10846a361472/page.htm' target='_blank' title='学院举行“厚植爱国情怀，砥砺强国之志”主题升旗仪式'>学院举行“厚植爱国情怀，砥砺强国之志”主...</a><i>12-23</i></li>
                        
                        <li><a href='/2024/1223/c10846a361470/page.htm' target='_blank' title='六WEI空间 | 包饺话冬至，合乐共此“食”'>六WEI空间 | 包饺话冬至，合乐共此“食”</a><i>12-23</i></li>
                        
                        <li><a href='/2024/1223/c10846a361467/page.htm' target='_blank' title='学院党委举行新党员入党宣誓暨青年党员重温入党誓词仪式'>学院党委举行新党员入党宣誓暨青年党员重温...</a><i>12-23</i></li>
                        
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
                
                <li><a href='/2024/1218/c10847a361184/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>12-18</i></li>
                
                <li><a href='/2024/1216/c10847a360991/page.htm' target='_blank' title='2023-2024学年教学优秀奖拟推荐名单公示'>2023-2024学年教学优秀奖拟推荐名单公示</a><i>12-16</i></li>
                
                <li><a href='/2024/1213/c10847a360780/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>12-13</i></li>
                
                <li><a href='/2024/1129/c10847a359561/page.htm' target='_blank' title='计算机科学与技术学院/软件学院2024-2025-2学期本科生教材选用公示'>计算机科学与技术学院/软件学院2024-2025-2...</a><i>11-29</i></li>
                
                <li><a href='/2024/1122/c10847a358972/page.htm' target='_blank' title='发展党员公示'>发展党员公示</a><i>11-25</i></li>
                
                <li><a href='/2024/1119/c10847a358735/page.htm' target='_blank' title='任前公示'>任前公示</a><i>11-19</i></li>
                
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
                
                <li><a href="/2024/1218/c10853a361184/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>12.18</span></li>
                
                <li><a href="/2024/1213/c10853a360780/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>12.13</span></li>
                
                <li><a href="/2024/1122/c10853a358972/page.htm" target="_blank"><i>></i>发展党员公示</a><span>11.25</span></li>
                
                <li><a href="/2024/1119/c10853a358731/page.htm" target="_blank"><i>></i>任前公示</a><span>11.19</span></li>
                
                <li><a href="/2024/0919/c10853a353834/page.htm" target="_blank"><i>></i>考察预告</a><span>09.19</span></li>
                
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
                
                <li><a href="/2024/1211/c10848a360585/page.htm" target="_blank"><i>></i>2024年度教职工考核工作安排</a><span>12.11</span></li>
                
                <li><a href="/2023/1214/c10848a360361/page.htm" target="_blank"><i>></i>2023年度教职工考核工作安排</a><span>12.14</span></li>
                
                <li><a href="/2023/0302/c10848a303998/page.htm" target="_blank"><i>></i>计算机科学与技术学院/人工智能学院/软件学...</a><span>03.02</span></li>
                
                <li><a href="/2022/1210/c10848a300564/page.htm" target="_blank"><i>></i>2022年度教职工考核工作安排</a><span>12.10</span></li>
                
                <li><a href="/2022/0922/c10848a293100/page.htm" target="_blank"><i>></i>2022年拟申报初级专业技术职务人员公示</a><span>09.22</span></li>
                
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
                
                <li><a href="/2024/1202/c10850a359812/page.htm" target="_blank"><i>></i>【毕业设计】2025届本科毕业设计工作通知</a><span>12.02</span></li>
                
                <li><a href="/2024/1012/c10850a355614/page.htm" target="_blank"><i>></i>【创新班专业分流】关于公布计算机科学与技...</a><span>10.12</span></li>
                
                <li><a href="/2024/1009/c10850a355263/page.htm" target="_blank"><i>></i>【创新班考核】2023-2024 学年 2021、2022...</a><span>10.09</span></li>
                
                <li><a href="/2024/0925/c10850a354395/page.htm" target="_blank"><i>></i>【创新班考核】关于开展2021级、2022级、20...</a><span>09.25</span></li>
                
                <li><a href="/2024/0916/c10850a353544/page.htm" target="_blank"><i>></i>【推免工作】2025届本科毕业生拟推荐免试攻...</a><span>09.16</span></li>
                
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
                
                <li><a href="/2025/0103/c10851a362183/page.htm" target="_blank"><i>></i><font style='color:#0A0300;'>【博士招生】2025年招收博士生申请考核制/...</font></a><span>01.03</span></li>
                
                <li><a href="/2024/1227/c10851a361749/page.htm" target="_blank"><i>></i>【博士答辩】杜祥通 博士答辩公告</a><span>12.27</span></li>
                
                <li><a href="/2024/1211/c10851a360543/page.htm" target="_blank"><i>></i>【奖学金】2023级硕士研究生学业奖学金动态...</a><span>12.11</span></li>
                
                <li><a href="/2024/1209/c10851a360309/page.htm" target="_blank"><i>></i>【导师】2024年第二批次硕导选聘公示</a><span>12.09</span></li>
                
                <li><a href="/2024/1205/c10851a360123/page.htm" target="_blank"><i>></i>【博士答辩】王善敏 博士答辩公告</a><span>12.05</span></li>
                
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
                <div class="date">2024-12-16</div>
              </div>
              <div class="title"> <a href='/2024/1213/c11617a360765/page.htm' target='_blank' title='【石榴大讲堂】不均衡学习的泛在适配'>【石榴大讲堂】不均衡学习的泛在适配</a>
                <p class="p1">题目：不均衡学习的泛在适配</p>
                <p class="p2">报告人：姚江超</p>
                <p class="p3">时间： <span class="t-time">2024-12-16</span> 下午3:00</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-12-16</div>
              </div>
              <div class="title"> <a href='/2024/1213/c11617a360764/page.htm' target='_blank' title='【石榴大讲堂】黑盒机器教学泛化理论：基础模型解析理论框架'>【石榴大讲堂】黑盒机器教学泛化理论：基础...</a>
                <p class="p1">题目：黑盒机器教学泛化理论：基础模型解析理论框架</p>
                <p class="p2">报告人：曹晓锋</p>
                <p class="p3">时间： <span class="t-time">2024-12-16</span> 下午2:00</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-12-15</div>
              </div>
              <div class="title"> <a href='/2024/1212/c11617a360653/page.htm' target='_blank' title='【石榴大讲堂】新AI时代的编译系统研究'>【石榴大讲堂】新AI时代的编译系统研究</a>
                <p class="p1">题目：新AI时代的编译系统研究</p>
                <p class="p2">报告人：江贺</p>
                <p class="p3">时间： <span class="t-time">2024-12-15</span> 上午10:00</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-12-10</div>
              </div>
              <div class="title"> <a href='/2024/1209/c11617a360306/page.htm' target='_blank' title='【石榴大讲堂】网络安全协议自动形式化验证系统'>【石榴大讲堂】网络安全协议自动形式化验证...</a>
                <p class="p1">题目：网络安全协议自动形式化验证系统</p>
                <p class="p2">报告人：熊焰</p>
                <p class="p3">时间： <span class="t-time">2024-12-10</span> 下午2:30</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-11-27</div>
              </div>
              <div class="title"> <a href='/2024/1126/c11617a359274/page.htm' target='_blank' title='【石榴大讲堂】Towards Time Series Foundation Models'>【石榴大讲堂】Towards Time Series Founda...</a>
                <p class="p1">题目：Towards Time Series Foundation Models</p>
                <p class="p2">报告人：吴敏</p>
                <p class="p3">时间： <span class="t-time">2024-11-27</span> 下午16:00</p>
                <p class="p4">地点：计算机学院楼505会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-11-25</div>
              </div>
              <div class="title"> <a href='/2024/1122/c11617a358958/page.htm' target='_blank' title='【石榴大讲堂】大图数据社区搜索的基础模型与方法'>【石榴大讲堂】大图数据社区搜索的基础模型...</a>
                <p class="p1">题目：大图数据社区搜索的基础模型与方法</p>
                <p class="p2">报告人：王国仁</p>
                <p class="p3">时间： <span class="t-time">2024-11-25</span> 上午10:00</p>
                <p class="p4">地点：计算机学院楼515会议室</p>
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
              <p>版权所有：南京航空航天大学计算机科学与技术学院/软件学院 ALL RIGHTS RESERVED 苏ICP备05070685号 <a href="http://site.nuaa.edu.cn/" target="_blank">后台管理</a> <a href="mailto:njliujr@nuaa.edu.cn" target="_blank"><img src="/_upload/tpl/04/02/1026/template1026/images/email.png"> 书记信箱</a> <a href="mailto:huangsj@nuaa.edu.cn" target="_blank"><img src="/_upload/tpl/04/02/1026/template1026/images/email.png"> 院长信箱</a></p>
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
