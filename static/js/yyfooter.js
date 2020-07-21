(function(){
    var scripts = document.getElementsByTagName('script');
    var src = scripts[scripts.length - 1].src;
    var arg = src.indexOf('?') !== -1 ? src.split('?').pop() : '';
    var settings = {};
    arg.replace(/(\w+)(?:=([^&]*))?/g, function(a, key, value) {
        settings[key] = value;
    });
    var thisYear = new Date().getFullYear();
    var className = (settings['type'] === 'mobile')?'yy-clean-footer ext-mobile':'yy-clean-footer';
    var brTag = (settings['type'] === 'mobile')?'<br>':'';
    var tpl = ['<div class="'+ className +'">',
    '©&nbsp;2005-'+ thisYear +'&nbsp;广州华多网络科技有限公司&nbsp;'+brTag+'版权所有&nbsp; <a href="http://www.beian.miit.gov.cn/" target="_blank">粤ICP备09075143号-13</a> &nbsp;<a href="http://bi2.duowan.com/app/index.php?r=h5page/copyright" target="_blank">版权保护投诉指引</a>',
            '</div>'].join('');
    var style = '<style>.yy-clean-footer {font-size: 12px;line-height: 18px;color: #333;text-align: center; padding: 28px 0;}</style>';

    document.write(style + tpl);
})()
