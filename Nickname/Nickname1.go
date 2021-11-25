/* Writed By iTz_Arshia */
/* Only works in #Nickname Channel */
/* Trigger Type : Regax */
/* Trigger : \A */

{{if hasRoleID "RoleID"}} /*inja ID Role OWNER ro mizarim ke Check kone agar Ersal konande Message On Role ro Dasht barash Embed paiin ro befreste*/
{{$y := sendMessageRetID nil (complexMessage
 "content" .User.Mention 
 "embed" (cembed "description" "**Man Nemitonam Nick Name Shoma ro Taghir bedam **" "color" 16776960))}}
{{deleteMessage nil $y 10}} /*inja bad az 10 Sanie payam ro pak mikone*/

{{else if lt (len .Message.Content) 33}} /*va in khat yani agar mohtavaye Content az Limitemon rad nashod in Etefagh biofe !!*/
{{editNickname .Message.Content}} /*NickName User be Mohtavaye Message esh taghir kone*/
{{$x := sendMessageRetID nil (complexMessage
 "content" .User.Mention
 "embed" (cembed "description" (print "**Nickname shoma az **__**" (or .Member.Nick .User.Username) "**__ **be** __**" .Message.Content "**__ **taghir kard**") "color" 65280))}}
{{deleteMessage nil $x 10}}

{{else}} /*va inja ham agar az in halat kharej bod yani bish az 32 Letter bood*/
{{$z := sendMessageRetID nil (complexMessage
 "content" .User.Mention
 "embed" (cembed "description" "**Nemitoni Toye Nickname az bishtar az __32__ harf estefade koni**" "color" 16776960))}}
{{deleteMessage nil $z 10}}
{{end}}
