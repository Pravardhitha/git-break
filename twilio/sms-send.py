from twilio.rest import Client
from creds import account_sid, auth_token, sender, toWhomst

client = Client(account_sid, auth_token)

sendsms = lambda x, body, z : client.messages.create(from_=x, body=body, to=z)

if __name__ == "__main__":
    message = ""
    sendsms(sender,message,toWhomst)