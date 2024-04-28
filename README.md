How does this profanity filter work?
This filter is to be used as an API. Imagine you have a chat application and you want to be able to make sure that in realtime you can check a message for any profanity. All you would need to do with this is pass the message via the body it will check for any banned words and then it will return the censored message back to you.


/v1/profanityfilter/{ID}
HTTP GET
ID (Required) - this would be the api token that is associated with the account. In this version it is limited with 10 calls. Returns the censored version (bad words replaced with asterisks) of any given text and whether the text contains profanity.

Parameters
text (required) - input text. Maximum 1000 characters.

Body
Message (required) - message that you would like to check for any banned words.

Sample Request URL Live Demo
https://www.lgmztech.com/v1/profanityfilter/{{UUID}}

Your Token ID: {{UUID}} - Max 10 Tries

You damn hooligan!
 Send Request

Sample Response
                
{"has-profanity":true,"modified-message":"You **** hooligan!","original-message":"You damn hooligan!"}