/*
 * Copyright 2025 coze-dev Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import type {
  IPlugin,
  IHooks,
  IPromptsHookParams,
} from 'rush-init-project-plugin';
// FIXME:
// 按照 https://github.com/bytemate/rush-plugins/blob/main/rush-plugins/rush-init-project-plugin/docs/init_project_configuration.md
// 一文的指引，无法正确 resolve 到对应模块，暂时没找到解决方案，故此处先用相对路径引用
// 未来需要调整为正常的 node_modules 引用方式
import { createLog } from '../../autoinstallers/plugins/node_modules/rush-init-project-plugin';
import { exec } from './utils';

export default class SetDefaultAuthorPlugin implements IPlugin {
  private readonly logger = createLog({
    prefix: SetDefaultAuthorPlugin.name,
  });

  apply(hooks: IHooks): void {
    hooks.prompts.tap(
      SetDefaultAuthorPlugin.name,
      async (prompts: IPromptsHookParams) => {
        const prompAuthorEmail = prompts.promptQueue.find(
          r => r.name === 'authorName',
        );
        if (prompAuthorEmail) {
          const userEmail = String(
            await exec(this.logger, 'git', ['config', '--get', 'user.email']),
          );
          Object.assign(prompAuthorEmail, {
            default() {
              return userEmail;
            },
            validate(author: string) {
              return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(author);
            },
          });

          hooks.answers.tap('authorPrefix', answers => {
            answers.authorPrefix = userEmail.split('@')?.[0] ?? '';
          });
        }
      },
    );
  }
}
