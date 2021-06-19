pattern = /^((https|retty):\/\/)?(reserve\.)?retty\.(me|co\.th|world)\/(purpose|area|category)(\/PRE\d{1}\d{1})?(\/ARE\d{1,})?(\/SUB\d{1,})?(\/STAN\d{1,})?(\/LND\d{1,})?(\/LCAT\d{1,})?(\/\d{12})?(\/)?$/
File.foreach('urls.txt') { |url| print url.chomp.chomp("/")[-1] if pattern.match? url.chomp }

