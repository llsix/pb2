print("开始打包")
local getTime = os.date("%Y%m%d%H%M%S")
local bv = gocom.buildv("bv")
local goenv = io.popen('go env GOARCH'):read("*all")
print("版本号",bv,"发布时间",getTime,"目标架构",goenv)
local pop = io.popen('go env GOARCH')
local goenv = pop:read("*all")
pop:close()

outfile = "pb2.exe"
packr_buildpath = "../../pb2/"

if string.find(goenv,"arm")~=nil then
	local ldflagv=string.format("-w -s -X 'main.PublishDate=%v' -X 'main.MinVersion=%v'",getTime, bv)

	_,er = gocom.exec({"go", "build", "-ldflags", ldflagv, "-o", "./build/out/"..outfile, "-tags=client"})
	
	if #er==0 then
		print("upx 开始压缩")
		out,er=gocom.exec({"upx", "-9", "./build/out/"..outfile})
		print(out,er)
	else 
		print(er)
	end
elseif string.find(goenv,"amd")~=nil then
--	_,er=gocom.exec({"go_packr2", "clean"})
--	_,er=gocom.exec({"go_packr2", "build", packr_buildpath })
	local ldflagv=string.format("-w -s -X 'main.PublishDate=%v' -X 'main.MinVersion=%v'",getTime, bv)
--	local ldflagv=string.format("-X 'main.PublishDate=%v' -X 'main.MinVersion=%v'",getTime, bv)

	_,er=gocom.exec({"go", "build", "-ldflags", ldflagv, "-o", "./build/out/"..outfile, "-tags=!client"})
	if #er==0 then
		print("upx 开始压缩")
		out,er=gocom.exec({"upx", "-9", "./build/out/"..outfile})
		print(out,er)
	else
		print(er)
	end
end


