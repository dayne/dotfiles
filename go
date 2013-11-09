#!/usr/bin/env ruby

def sym_link_it(rel_file,rel_target)
  # full path to file
  file = File.join( ENV['HOME'], '.dotfiles', rel_file )
  target = File.join( ENV['HOME'], rel_target )
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


sym_link_it( 'vim/vimrc.symlink' , '.vimrc' )
sym_link_it( 'irbrc.symlink' ,  '.irbrc'  )

system("bundle")

