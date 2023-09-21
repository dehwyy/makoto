export default {
    "name": "SignUp",
    "kind": "HoudiniMutation",
    "hash": "d15773e558c0d7de656867e3567af1b88b693fbc6ee9600460dd950036cb3606",

    "raw": `mutation SignUp($username: String!, $password: String!, $question: String!, $answer: String!) {
  signUp(
    input: {username: $username, password: $password, question: $question, answer: $answer}
  ) {
    userId
    tokens {
      access_token
      refresh_token
    }
  }
}
`,

    "rootType": "Mutation",

    "selection": {
        "fields": {
            "signUp": {
                "type": "UserAuthResponse",
                "keyRaw": "signUp(input: {username: $username, password: $password, question: $question, answer: $answer})",

                "selection": {
                    "fields": {
                        "userId": {
                            "type": "ID",
                            "keyRaw": "userId",
                            "visible": true
                        },

                        "tokens": {
                            "type": "Tokens",
                            "keyRaw": "tokens",

                            "selection": {
                                "fields": {
                                    "access_token": {
                                        "type": "String",
                                        "keyRaw": "access_token",
                                        "visible": true
                                    },

                                    "refresh_token": {
                                        "type": "String",
                                        "keyRaw": "refresh_token",
                                        "visible": true
                                    }
                                }
                            },

                            "visible": true
                        }
                    }
                },

                "visible": true
            }
        }
    },

    "pluginData": {
        "houdini-svelte": {}
    },

    "input": {
        "fields": {
            "username": "String",
            "password": "String",
            "question": "String",
            "answer": "String"
        },

        "types": {}
    }
};

"HoudiniHash=fcbf307d9ec847dcc3207c72b386bb383951c3d1fcde8770b362e999168404ea";