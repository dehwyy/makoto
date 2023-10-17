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
  GOOGLE_CLIENT_ID: str({ default: '' }),
  GOOGLE_CLIENT_SECRET: str({ default: '' }),
  GOOGLE_REDIRECT_URL: str({ default: '' }),

  // mode
	NODE_ENV: str({ choices: ['development', 'production'], default: 'development' }),
});

export enum SERVER_PORTS {
  AUTH = 5001,
  HASHMAP
}

export enum PORTS {
  HOME = 3000,
  AUTH,
  HASHMAP
}
