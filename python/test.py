import mod_config
mod_config.iniConfig("python.conf")
text1 = mod_config.getConfig("getnovelcontent", "text1")
print text1
