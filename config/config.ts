import {parse} from "yaml"
import fs from "fs"

interface Config {
  ports: {
    gateway: string
    auth: string
  }
  frontend_ports: {
    auth: string
  }
  oauth2: {
    google: {
      client_id: string
      secret: string
      redirect_url: string
    }
  }
}


export const config = (path_to_cfg: string) => parse(fs.readFileSync(path_to_cfg + "/config.yaml", "utf-8")) as Config
