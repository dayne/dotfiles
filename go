#!/usr/bin/env ruby

def sym_link_it(file,target)
	backup = target+'.dib' # d init backup
	if File.symlink?(target) 
		puts "skipping #{file} as the target already a symlink: #{target}"
		return false
	end
	if File.exists?(backup)
		puts "skipping #{file} as d.init backup exists already: #{backup}"
		return false
	end
	if File.exists?(target)
		puts "backing up #{File.basename(target)} as #{backup}"
		File.mv(target, backup)
	end

	puts "creating symlink: #{target} -> #{file}"
	File.symlink(file,target)
end

sym_link_it( File.join( Dir.pwd,'vim/vimrc.symlink' ), File.join( ENV['HOME'], '.vimrc' ) )

