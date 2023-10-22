import { resolve } from 'path';
import * as dotenv from 'dotenv';
import { cleanEnv, str } from 'envalid';

try {
	dotenv.config({ path: resolve(process.cwd(), '../../.env') });
} catch {
  console.log('No .env file found');
}

export const config = cleanEnv(process.env, {
  // OAauth2
  // google
  GOOGLE_CLIENT_ID: str({ default: '' }),
  GOOGLE_CLIENT_SECRET: str({ default: '' }),
  GOOGLE_REDIRECT_URL: str({ default: '' }),
  // github
  GITHUB_CLIENT_ID: str({default: ""}),
  GITHUB_CLIENT_SECRET: str({default: ""}),
  GITHUB_REDIRECT_URL: str({default: ""}),
  GITHUB_STATE: str({default: ""}),
  // discord
  DISCORD_CLIENT_ID:str({default: ""}),
  DISCORD_CLIENT_SECRET:str({default: ""}),
  DISCORD_REDIRECT_URL:str({default: ""}),
  DISCORD_STATE:str({default: ""}),


  // mode
	NODE_ENV: str({ choices: ['development', 'production'], default: 'development' }),
});

export enum SERVER_PORTS {
  MAIN = 5000,
  AUTH,
  HASHMAP
}

export enum PORTS {
  MAIN = 3000,
  AUTH,
  HASHMAP,
  DISCORD
}
