import { BaseListResp } from '/@/api/model/baseModel';

/**
 *  @description: Audit info response
 */
export interface AuditInfo {
  objectName?: string;
  actionName?: string;
  changedData?: string;
  createdBy?: string;
}

/**
 *  @description: Audit list response
 */

export type AuditListResp = BaseListResp<AuditInfo>;
