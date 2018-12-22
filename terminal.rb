def init
  puts("###################################################")
  puts("######### Otimizando comandos do terminal #########")
  puts("###################################################")
  puts(" MENU ")
  puts(" 1 - Dando start no PostgreSQL ")
  puts(" 2 - Dando stop no PostgreSQL ")
  puts(" 0 - Sair ")

  # Capiturando o
  command = gets()

  case command.to_i
  when 0
    puts "Tchau... :)"
  when 1
    exec 'pg_ctl -D /usr/local/var/postgres start'
  when 2
    exec 'pg_ctl -D /usr/local/var/postgres stop'
  else
    puts "Opção não catalogada"
    sleep 1
    init
  end
end

init