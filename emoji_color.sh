
# Screen output nice colors
declare -A colors
# Reset
colors[Color_Off]='\033[0m'       # Text Reset

# Regular Colors
colors[Red]='\033[0;31m'          # Red
colors[Green]='\033[0;32m'        # Green
colors[Yellow]='\033[0;33m'       # Yellow
colors[Blue]='\033[0;34m'         # Blue
colors[Cyan]='\033[0;36m'         # Cyan
colors[White]='\033[0;37m'        # White

# Bold
colors[BRed]='\033[1;31m'         # Red
colors[BGreen]='\033[1;32m'       # Green
colors[BYellow]='\033[1;33m'      # Yellow
colors[BBlue]='\033[1;34m'        # Blue
colors[BPurple]='\033[1;35m'      # Purple
colors[BCyan]='\033[1;36m'        # Cyan
colors[BWhite]='\033[1;37m'       # White

# Underline
colors[URed]='\033[4;31m'         # Red
colors[UGreen]='\033[4;32m'       # Green
colors[UYellow]='\033[4;33m'      # Yellow
colors[UBlue]='\033[4;34m'        # Blue
colors[UCyan]='\033[4;36m'        # Cyan
colors[UWhite]='\033[4;37m'       # White

# Background
colors[On_Black]='\033[40m'       # Black
colors[On_Red]='\033[41m'         # Red
colors[On_Green]='\033[42m'       # Green
colors[On_Yellow]='\033[30;43m'   # Yellow
colors[On_Blue]='\033[44m'        # Blue
colors[On_Purple]='\033[45m'      # Purple
colors[On_Cyan]='\033[30;46m'     # Cyan

# High Intensity
colors[IRed]='\033[0;91m'         # Red
colors[IGreen]='\033[0;92m'       # Green
colors[IYellow]='\033[0;93m'      # Yellow
colors[IBlue]='\033[0;94m'        # Blue
colors[IPurple]='\033[0;95m'      # Purple
colors[ICyan]='\033[0;96m'        # Cyan
colors[IWhite]='\033[0;97m'       # White

# Bold High Intensity
colors[BIRed]='\033[1;91m'        # Red
colors[BIGreen]='\033[1;92m'      # Green
colors[BIYellow]='\033[1;93m'     # Yellow
colors[BIBlue]='\033[1;94m'       # Blue
colors[BIPurple]='\033[1;95m'     # Purple
colors[BICyan]='\033[30;1;96m'    # Cyan
colors[BIWhite]='\033[1;97m'      # White

# High Intensity backgrounds
colors[On_IBlack]='\033[0;100m'   # Black
colors[On_IRed]='\033[0;101m'     # Red
colors[On_IGreen]='\033[0;102m'   # Green
colors[On_IYellow]='\033[0;30;103m' # Yellow
colors[On_IBlue]='\033[0;30;104m'    # Blue
colors[On_IPurple]='\033[0;105m'  # Purple
colors[On_ICyan]='\033[0;30;106m'    # Cyan

# Emoji
declare -A emoji
emoji[Robot]='\U1F916'           #Robot
emoji[Poo]='\U1F4A9'             #Poo
emoji[Alien]='\U1F47E'           #Alien
emoji[Eyes]='\U1F440'            #Eyes
emoji[Whale]='\U1F40B'           #Whale
emoji[Fire]='\U1F525'            #Fire
emoji[Danger]='\U26A1'           #Danger
emoji[HammerRench]='\U1F6E0'     #Hammer and Rench
emoji[Bomb]='\U1F4A3'            #Bomb
emoji[Gear]='\U2699'             #Gear
emoji[NoEntry]='\U26D4'          #No Entry
emoji[Radioactive]='\U2622'      #Radioactive
emoji[Biohazard]='\U2623'        #Biohazard
emoji[Broom]='\U1F9F9'           #Broom
emoji[Toilet]='\U1F6BD'          #Toilet
emoji[Locked]='\U1F512'          #Locked
emoji[Construction]='\U1F6A7'    #Construction
emoji[Cactus]='\U1F335'          #Cactus
emoji[SkullBones]='\U2620'       #Skull and Bones


echocolor(){
    # $1 - the text, $2 - color, $3 - emoji to display
    color=${colors[White]}
    icon=${emoji[Robot]}
    if [[ ! -z "$2" ]]; then
    	color=${colors[$2]}
    fi
    if [[ ! -z "$3"  ]]; then
		icon=${emoji[$3]}
    fi

	echo -e "$icon $color $1 ${colors[Color_Off]}"
}

# end of screen color

test_color(){
  for i in "${!colors[@]}"
  do
  	echo -e "${colors[Color_Off]}$i = ${colors[$i]}Color test${colors[White]}"
  done	
}

test_emoji(){
  for i in "${!emoji[@]}"
  do
  	echo -e "$i = ${emoji[$i]}"
  done	
}


colors_formatting(){
	# This program is free software. It comes without any warranty, to
	# the extent permitted by applicable law. You can redistribute it
	# and/or modify it under the terms of the Do What The Fuck You Want
	# To Public License, Version 2, as published by Sam Hocevar. See
	# http://sam.zoy.org/wtfpl/COPYING for more details.
	 
	#Background
	for clbg in {40..47} {100..107} 49 ; do
		#Foreground
		for clfg in {30..37} {90..97} 39 ; do
			#Formatting
			for attr in 0 1 2 4 5 7 ; do
				#Print the result
				echo -en "\e[${attr};${clbg};${clfg}m ^[${attr};${clbg};${clfg}m \e[0m"
			done
			echo #Newline
		done
	done
}

test256colors(){
	# This program is free software. It comes without any warranty, to
	# the extent permitted by applicable law. You can redistribute it
	# and/or modify it under the terms of the Do What The Fuck You Want
	# To Public License, Version 2, as published by Sam Hocevar. See
	# http://sam.zoy.org/wtfpl/COPYING for more details.
	 
	for fgbg in 38 48 ; do # Foreground / Background
	    for color in {0..255} ; do # Colors
	        # Display the color
	        printf "\e[${fgbg};5;%sm  %3s  \e[0m" $color $color
	        # Display 6 colors per lines
	        if [ $((($color + 1) % 6)) == 4 ] ; then
	            echo # New line
	        fi
	    done
	    echo # New line
	done
}
