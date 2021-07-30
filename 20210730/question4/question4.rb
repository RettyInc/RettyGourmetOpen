#!/usr/bin/env ruby

# 画像ファイルの正しい種別を判定する
# JPEG/PNG/GIF/その他(OTHER)

def image_type(filepath)
    header = ""
    File.open(filepath) do |f|
        header = f.read(8)
    end

    return 'OTHER' if header.nil?

    type = if header[0, 2].unpack('H*') == %w(ffd8) then
        'JPEG'
    elsif header[0, 4].unpack('H*') == %w(89504e47) then
        'PNG'
    elsif header[0, 4].unpack('H*') == %w(47494638) then
        'GIF'
    else
        'OTHER'
    end

    return type
end

unless ARGV[0]
    puts "画像ファイルへのパスを引数で与えてください"
    exit(1)
end

puts image_type(ARGV[0])