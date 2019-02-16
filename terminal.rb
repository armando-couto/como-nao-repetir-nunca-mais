def init
  puts("###################################################")
  puts("######### Otimizando comandos do terminal #########")
  puts("###################################################")
  puts(" MENU ")
  puts(" 1 - PostgreSQL - Start ")
  puts(" 2 - PostgreSQL - Stop ")
  puts(" 3 - Redis - Start ")
  puts(" 4 - Mongo DB - Start ")
  puts(" 5 - Apache - Start ")
  puts(" 6 - Apache - Stop ")
  puts(" 7 - MySQL - Start ")
  puts(" 8 - MySQL - Stop ")
  puts(" 0 - Sair ")

  # Capiturando o
  command = gets()

  case command.to_i
  when 0
    puts "Tchau... :)"
  when 1
    exec 'pg_ctl -D /usr/local/var/postgres start &'
  when 2
    exec 'pg_ctl -D /usr/local/var/postgres stop &'
  when 3
    exec 'redis-server /usr/local/etc/redis.conf'
  when 4
    exec 'mongod'
  when 5
    exec 'sudo /usr/sbin/apachectl start'
  when 6
    exec 'sudo /usr/sbin/apachectl stop'
  when 7
    exec 'brew services start mysql'
  when 8
    exec 'brew services stop mysql'
  else
    puts "Opção não catalogada"
    sleep 1
    init
  end
end

init