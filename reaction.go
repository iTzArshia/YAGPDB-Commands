/* Writed By iTz_Arshia */
/* Only works in #Nickname Channel */
/* Trigger Type : Reaction Add */

{{deleteMessageReaction .Channel.ID .Message.ID .User.ID "Emoji"}} /*inja React Member ro pak mikone*/
{{editNickname .User.Username}} /*inja ham Nickname Member ro be usernamesh edit mikone*/
 
{{$cooldown := 3600}} /*Time out 1 saate ke yani agar ye nafar bezane ro react va nicknamesh reseet beshe ta 1 saat bad harcheghadram bezane ro react embed nemifreste*/
{{if $cd := dbGet .CCID "cooldown" }}
{{else}}
{{dbSetExpire .CCID "cooldown" "cooldown" $cooldown}}
 
{{$x := sendMessageRetID nil (complexMessage
 "content" .User.Mention
 "embed" (cembed "description" (print "**Nickname shoma be ** __**" .User.Username "**__  **taghir kard**") "color" 16776960))}}
 
{{deleteMessage nil $x 10}} /*inam ke hamontor ke ghablan goftam bade 10 sanie payamo pak mikone*/
{{end}}
 

