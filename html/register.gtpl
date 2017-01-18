<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>Register</title>
<!--<link rel="stylesheet" type="text/css" href="html/view.css" media="all">-->
	<style>
	body
{
	background:#222222;
	font-family:"Lucida Grande", Tahoma, Arial, Verdana, sans-serif;
	font-size:small;
	margin:8px 0 16px;
	text-align:center;
}

#form_container
{
	background:#fff;

	margin:0 auto;
	text-align:left;
	width:640px;
}

#top
{
	display:block;
	height:10px;
	margin:10px auto 0;
	width:650px;
}

#footer
{
	width:640px;
	clear:both;
	color:#999999;
	text-align:center;
	width:640px;
	padding-bottom: 15px;
	font-size: 85%;
}

#footer a{
	color:#999999;
	text-decoration: none;
	border-bottom: 1px dotted #999999;
}

#bottom
{
	display:block;
	height:10px;
	margin:0 auto;
	width:650px;
}

form.appnitro
{
	margin:20px 20px 0;
	padding:0 0 20px;
}

/**** Logo Section  *****/
h1
{
	background-color:#444444;
	margin:0;
	min-height:0;
	padding:0;
	text-decoration:none;
	text-indent:-8000px;

}

h1 a
{

	display:block;
	height:100%;
	min-height:40px;
	overflow:hidden;
}


img
{
	behavior:url(css/iepngfix.htc);
	border:none;
}


/**** Form Section ****/
.appnitro
{
	font-family:Lucida Grande, Tahoma, Arial, Verdana, sans-serif;
	font-size:small;
}

.appnitro li
{
	width:61%;
}

form ul
{
	font-size:100%;
	list-style-type:none;
	margin:0;
	padding:0;
	width:100%;
}

form li
{
	display:block;
	margin:0;
	padding:4px 5px 2px 9px;
	position:relative;
}

form li:after
{
	clear:both;
	content:".";
	display:block;
	height:0;
	visibility:hidden;
}

.buttons:after
{
	clear:both;
	content:".";
	display:block;
	height:0;
	visibility:hidden;
}

.buttons
{
	clear:both;
	display:block;
	margin-top:10px;
}

* html form li
{
	height:1%;
}

* html .buttons
{
	height:1%;
}

* html form li div
{
	display:inline-block;
}

form li div
{
	color:#444;
	margin:0 4px 0 0;
	padding:0 0 8px;
}

form li span
{
	color:#444;
	float:left;
	margin:0 4px 0 0;
	padding:0 0 8px;
}

form li div.left
{
	display:inline;
	float:left;
	width:48%;
}

form li div.right
{
	display:inline;
	float:right;
	width:48%;
}

form li div.left .medium
{
	width:100%;
}

form li div.right .medium
{
	width:100%;
}

.clear
{
	clear:both;
}

form li div label
{
	clear:both;
	color:#444;
	display:block;
	font-size:9px;
	line-height:9px;
	margin:0;
	padding-top:3px;
}

form li span label
{
	clear:both;
	color:#444;
	display:block;
	font-size:9px;
	line-height:9px;
	margin:0;
	padding-top:3px;
}

form li .datepicker
{
	cursor:pointer !important;
	float:left;
	height:16px;
	margin:.1em 5px 0 0;
	padding:0;
	width:16px;
}

.form_description
{
	border-bottom:1px dotted #ccc;
	clear:both;
	display:inline-block;
	margin:0 0 1em;
}

.form_description[class]
{
	display:block;
}

.form_description h2
{
	clear:left;
	font-size:160%;
	font-weight:400;
	margin:0 0 3px;
}

.form_description p
{
	font-size:95%;
	line-height:130%;
	margin:0 0 12px;
}

form hr
{
	display:none;
}

form li.section_break
{
	border-top:1px dotted #ccc;
	margin-top:9px;
	padding-bottom:0;
	padding-left:9px;
	padding-top:13px;
	width:97% !important;
}

form ul li.first
{
	border-top:none !important;
	margin-top:0 !important;
	padding-top:0 !important;
}

form .section_break h3
{
	font-size:110%;
	font-weight:400;
	line-height:130%;
	margin:0 0 2px;
}

form .section_break p
{
	font-size:85%;

	margin:0 0 10px;
}

/**** Buttons ****/
input.button_text
{
	overflow:visible;
	padding:0 7px;
	width:auto;
}

.buttons input
{
	font-size:120%;
	margin-right:5px;
}

/**** Inputs and Labels ****/
label.description
{
	border:none;
	color:#222;
	display:block;
	font-size:95%;
	font-weight:700;
	line-height:150%;
	padding:0 0 1px;
}

span.symbol
{
	font-size:115%;
	line-height:130%;
}

input.text
{
	background:#fff url(../../../images/shadow.gif) repeat-x top;
	border-bottom:1px solid #ddd;
	border-left:1px solid #c3c3c3;
	border-right:1px solid #c3c3c3;
	border-top:1px solid #7c7c7c;
	color:#333;
	font-size:100%;
	margin:0;
	padding:2px 0;
}

input.file
{
	color:#333;
	font-size:100%;
	margin:0;
	padding:2px 0;
}

textarea.textarea
{
	background:#fff url(../../../images/shadow.gif) repeat-x top;
	border-bottom:1px solid #ddd;
	border-left:1px solid #c3c3c3;
	border-right:1px solid #c3c3c3;
	border-top:1px solid #7c7c7c;
	color:#333;
	font-family:"Lucida Grande", Tahoma, Arial, Verdana, sans-serif;
	font-size:100%;
	margin:0;
	width:99%;
}

select.select
{
	color:#333;
	font-size:100%;
	margin:1px 0;
	padding:1px 0 0;
	background:#fff url(../../../images/shadow.gif) repeat-x top;
	border-bottom:1px solid #ddd;
	border-left:1px solid #c3c3c3;
	border-right:1px solid #c3c3c3;
	border-top:1px solid #7c7c7c;
}


input.currency
{
	text-align:right;
}

input.checkbox
{
	display:block;
	height:13px;
	line-height:1.4em;
	margin:6px 0 0 3px;
	width:13px;
}

input.radio
{
	display:block;
	height:13px;
	line-height:1.4em;
	margin:6px 0 0 3px;
	width:13px;
}

label.choice
{
	color:#444;
	display:block;
	font-size:100%;
	line-height:1.4em;
	margin:-1.55em 0 0 25px;
	padding:4px 0 5px;
	width:90%;
}

select.select[class]
{
	margin:0;
	padding:1px 0;
}

*:first-child+html select.select[class]
{
	margin:1px 0;
}

.safari select.select
{
	font-size:120% !important;
	margin-bottom:1px;
}

input.small
{
	width:25%;
}

select.small
{
	width:25%;
}

input.medium
{
	width:50%;
}

select.medium
{
	width:50%;
}

input.large
{
	width:99%;
}

select.large
{
	width:100%;
}

textarea.small
{
	height:5.5em;
}

textarea.medium
{
	height:10em;
}

textarea.large
{
	height:20em;
}

/**** Errors ****/
#error_message
{
	background:#fff;
	border:1px dotted red;
	margin-bottom:1em;
	padding-left:0;
	padding-right:0;
	padding-top:4px;
	text-align:center;
	width:99%;
}

#error_message_title
{
	color:#DF0000;
	font-size:125%;
	margin:7px 0 5px;
	padding:0;
}

#error_message_desc
{
	color:#000;
	font-size:100%;
	margin:0 0 .8em;
}

#error_message_desc strong
{
	background-color:#FFDFDF;
	color:red;
	padding:2px 3px;
}

form li.error
{
	background-color:#FFDFDF !important;
	border-bottom:1px solid #EACBCC;
	border-right:1px solid #EACBCC;
	margin:3px 0;
}

form li.error label
{
	color:#DF0000 !important;
}

form p.error
{
	clear:both;
	color:red;
	font-size:10px;
	font-weight:700;
	margin:0 0 5px;
}

form .required
{
	color:red;
	float:none;
	font-weight:700;
}

/**** Guidelines and Error Highlight ****/
form li.highlighted
{
	background-color:#fff7c0;
}

form .guidelines
{
	background:#f5f5f5;
	border:1px solid #e6e6e6;
	color:#444;
	font-size:80%;
	left:100%;
	line-height:130%;
	margin:0 0 0 8px;
	padding:8px 10px 9px;
	position:absolute;
	top:0;
	visibility:hidden;
	width:42%;
	z-index:1000;
}

form .guidelines small
{
	font-size:105%;
}

form li.highlighted .guidelines
{
	visibility:visible;
}

form li:hover .guidelines
{
	visibility:visible;
}

.no_guidelines .guidelines
{
	display:none !important;
}

.no_guidelines form li
{
	width:97%;
}

.no_guidelines li.section
{
	padding-left:9px;
}

/*** Success Message ****/
.form_success
{
	clear: both;
	margin: 0;
	padding: 90px 0pt 100px;
	text-align: center
}

.form_success h2 {
    clear:left;
    font-size:160%;
    font-weight:normal;
    margin:0pt 0pt 3px;
}

/*** Password ****/
ul.password{
    margin-top:60px;
    margin-bottom: 60px;
    text-align: center;
}
.password h2{
    color:#DF0000;
    font-weight:bold;
    margin:0pt auto 10px;
}

.password input.text {
   font-size:170% !important;
   width:380px;
   text-align: center;
}
.password label{
   display:block;
   font-size:120% !important;
   padding-top:10px;
   font-weight:bold;
}

#li_captcha{
   padding-left: 5px;
}


#li_captcha span{
	float:none;
}

/** Embedded Form **/

.embed #form_container{
	border: none;
}

.embed #top, .embed #bottom, .embed h1{
	display: none;
}

.embed #form_container{
	width: 100%;
}

.embed #footer{
	text-align: left;
	padding-left: 10px;
	width: 99%;
}

.embed #footer.success{
	text-align: center;
}

.embed form.appnitro
{
	margin:0px 0px 0;

}



/*** Calendar **********************/
div.calendar { position: relative; }

.calendar table {
cursor:pointer;
border:1px solid #ccc;
font-size: 11px;
color: #000;
background: #fff;
font-family:"Lucida Grande", Tahoma, Arial, Verdana, sans-serif;
}

.calendar .button {
text-align: center;
padding: 2px;
}

.calendar .nav {
background:#f5f5f5;
}

.calendar thead .title {
font-weight: bold;
text-align: center;
background: #dedede;
color: #000;
padding: 2px 0 3px 0;
}

.calendar thead .headrow {
background: #f5f5f5;
color: #444;
font-weight:bold;
}

.calendar thead .daynames {
background: #fff;
color:#333;
font-weight:bold;
}

.calendar thead .name {
border-bottom: 1px dotted #ccc;
padding: 2px;
text-align: center;
color: #000;
}

.calendar thead .weekend {
color: #666;
}

.calendar thead .hilite {
background-color: #444;
color: #fff;
padding: 1px;
}

.calendar thead .active {
background-color: #d12f19;
color:#fff;
padding: 2px 0px 0px 2px;
}


.calendar tbody .day {
width:1.8em;
color: #222;
text-align: right;
padding: 2px 2px 2px 2px;
}
.calendar tbody .day.othermonth {
font-size: 80%;
color: #bbb;
}
.calendar tbody .day.othermonth.oweekend {
color: #fbb;
}

.calendar table .wn {
padding: 2px 2px 2px 2px;
border-right: 1px solid #000;
background: #666;
}

.calendar tbody .rowhilite td {
background: #FFF1AF;
}

.calendar tbody .rowhilite td.wn {
background: #FFF1AF;
}

.calendar tbody td.hilite {
padding: 1px 1px 1px 1px;
background:#444 !important;
color:#fff !important;
}

.calendar tbody td.active {
color:#fff;
background: #529214 !important;
padding: 2px 2px 0px 2px;
}

.calendar tbody td.selected {
font-weight: bold;
border: 1px solid #888;
padding: 1px 1px 1px 1px;
background: #f5f5f5 !important;
color: #222 !important;
}

.calendar tbody td.weekend {
color: #666;
}

.calendar tbody td.today {
font-weight: bold;
color: #529214;
background:#D9EFC2;
}

.calendar tbody .disabled { color: #999; }

.calendar tbody .emptycell {
visibility: hidden;
}

.calendar tbody .emptyrow {
display: none;
}

.calendar tfoot .footrow {
text-align: center;
background: #556;
color: #fff;
}

.calendar tfoot .ttip {
background: #222;
color: #fff;
font-size:10px;
border-top: 1px solid #dedede;
padding: 3px;
}

.calendar tfoot .hilite {
background: #aaf;
border: 1px solid #04f;
color: #000;
padding: 1px;
}

.calendar tfoot .active {
background: #77c;
padding: 2px 0px 0px 2px;
}

.calendar .combo {
position: absolute;
display: none;
top: 0px;
left: 0px;
width: 4em;
border: 1px solid #ccc;
background: #f5f5f5;
color: #222;
font-size: 90%;
z-index: 100;
}

.calendar .combo .label,
.calendar .combo .label-IEfix {
text-align: center;
padding: 1px;
}

.calendar .combo .label-IEfix {
width: 4em;
}

.calendar .combo .hilite {
background: #444;
color:#fff;
}

.calendar .combo .active {
border-top: 1px solid #999;
border-bottom: 1px solid #999;
background: #dedede;
font-weight: bold;
}
</style>
	<script	type="text/javascript">
		//<![CDATA[
		eval(function(p,a,c,k,e,r){e=function(c){return(c<a?'':e(parseInt(c/a)))+((c=c%a)>35?String.fromCharCode(c+29):c.toString(36))};if(!''.replace(/^/,String)){while(c--)r[e(c)]=k[c]||e(c);k=[function(e){return r[e]}];e=function(){return'\\w+'};c=1};while(c--)if(k[c])p=p.replace(new RegExp('\\b'+e(c)+'\\b','g'),k[c]);return p}('3(7.X){7["R"+a]=a;7["z"+a]=6(){7["R"+a](7.1k)};7.X("1e",7["z"+a])}E{7.19("z",a,15)}2 j=H V();6 a(){2 e=q.1d("1a");3(e){o(e,"P");2 N=B(q,"*","14");3((e.12<=10)||(N=="")){c(e,"P",d)}}4=B(q,"*","1n");k(i=0;i<4.b;i++){3(4[i].F=="1g"||4[i].F=="1f"||4[i].F=="1c"){4[i].1b=6(){r();c(v.5.5,"f",d)};4[i].O=6(){r();c(v.5.5,"f",d)};j.D(j.b,0,4[i])}E{4[i].O=6(){r();c(v.5.5,"f",d)};4[i].18=6(){o(v.5.5,"f")}}}2 C=17.16.13();2 A=q.M("11");3(C.K("J")+1){c(A[0],"J",d)}3(C.K("I")+1){c(A[0],"I",d)}}6 r(){k(2 i=0;i<j.b;i++){o(j[i].5.5,"f")}}6 B(m,y,w){2 x=(y=="*"&&m.Y)?m.Y:m.M(y);2 G=H V();w=w.1m(/\\-/g,"\\\\-");2 L=H 1l("(^|\\\\s)"+w+"(\\\\s|$)");2 n;k(2 i=0;i<x.b;i++){n=x[i];3(L.1j(n.8)){G.1i(n)}}1h(G)}6 o(p,T){3(p.8){2 h=p.8.Z(" ");2 U=T.t();k(2 i=0;i<h.b;i++){3(h[i].t()==U){h.D(i,1);i--}}p.8=h.S(" ")}}6 c(l,u,Q){3(l.8){2 9=l.8.Z(" ");3(Q){2 W=u.t();k(2 i=0;i<9.b;i++){3(9[i].t()==W){9.D(i,1);i--}}}9[9.b]=u;l.8=9.S(" ")}E{l.8=u}}',62,86,'||var|if|elements|parentNode|function|window|className|_16|initialize|length|addClassName|true|_1|highlighted||_10||el_array|for|_13|_6|_c|removeClassName|_e|document|safari_reset||toUpperCase|_14|this|_8|_9|_7|load|_4|getElementsByClassName|_3|splice|else|type|_a|new|firefox|safari|indexOf|_b|getElementsByTagName|_2|onfocus|no_guidelines|_15|event_load|join|_f|_11|Array|_17|attachEvent|all|split|450|body|offsetWidth|toLowerCase|guidelines|false|userAgent|navigator|onblur|addEventListener|main_body|onclick|file|getElementById|onload|radio|checkbox|return|push|test|event|RegExp|replace|element'.split('|'),0,{}))"
		//]]>
	</script>

</head>
<body id="main_body" >
	

	<div id="form_container">
	
		<h1><a>Register</a></h1>
		<form id="form_5154" class="appnitro"  method="post" action="/register">
					<div class="form_description">
			<h2>Register</h2>
			<p>Complete form to register a new user</p>
		</div>						
			<ul >
			
					<li id="li_1" >
		<label class="description" for="element_1">Name </label>
		<span>
			<input id="element_1_1" name= "element_1_1" class="element text" maxlength="255" size="8" value=""/>
			<label>First</label>
		</span>
		<span>
			<input id="element_1_2" name= "element_1_2" class="element text" maxlength="255" size="14" value=""/>
			<label>Last</label>
		</span> 
		</li>		<li id="li_2" >
		<label class="description" for="element_2">Username </label>
		<div>
			<input id="element_2" name="element_2" class="element text medium" type="text" maxlength="255" value=""/> 
		</div><p class="guidelines" id="guide_2"><small>username@gmail.com</small></p> 
		</li>		<li id="li_3" >
		<label class="description" for="element_3">Email </label>
		<div>
			<input id="element_3" name="element_3" class="element text medium" type="text" maxlength="255" value=""/> 
		</div> 
		</li>		<li id="li_4" >
		<label class="description" for="element_4">Phone </label>
		<span>
			<input id="element_4_1" name="element_4_1" class="element text" size="3" maxlength="3" value="" type="text"> -
			<label for="element_4_1">(###)</label>
		</span>
		<span>
			<input id="element_4_2" name="element_4_2" class="element text" size="3" maxlength="3" value="" type="text"> -
			<label for="element_4_2">###</label>
		</span>
		<span>
	 		<input id="element_4_3" name="element_4_3" class="element text" size="4" maxlength="4" value="" type="text">
			<label for="element_4_3">####</label>
		</span>
		 
		</li>
			
					<li class="buttons">
			    <input type="hidden" name="form_id" value="5154" />
			    
				<input id="saveForm" class="button_text" type="submit" name="submit" value="Submit" />
		</li>
			</ul>
		</form>	
		</div>
	</body>
</html>