/**
 * Mini-app text inputs always produce strings. Coerce selected fields to the
 * JSON types the Go API expects before calling uni.request.
 */

export type CoerceKind = 'number' | 'boolean' | 'date' | 'string'

export type CoerceSchema = Record<string, CoerceKind>

function toNumber(value: unknown): number {
  if (typeof value === 'number' && Number.isFinite(value)) {
    return value
  }
  if (value === null || value === undefined || value === '') {
    return 0
  }
  const n = Number(String(value).trim())
  return Number.isFinite(n) ? n : 0
}

function toBoolean(value: unknown): boolean {
  if (typeof value === 'boolean') {
    return value
  }
  if (value === null || value === undefined || value === '') {
    return false
  }
  const s = String(value).trim().toLowerCase()
  return s === 'true' || s === '1' || s === 'yes'
}

/** YYYY-MM-DD / local datetime → RFC3339-ish string Go time.Time accepts. */
function toDate(value: unknown): string | null {
  if (value === null || value === undefined || value === '') {
    return null
  }
  const s = String(value).trim()
  if (!s) {
    return null
  }
  if (/^\d{4}-\d{2}-\d{2}$/.test(s)) {
    return `${s}T00:00:00Z`
  }
  // uni formatDate often emits "YYYY-MM-DDTHH:mm:ss" without zone
  if (/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$/.test(s)) {
    return `${s}Z`
  }
  return s
}

function toString(value: unknown): string {
  if (value === null || value === undefined) {
    return ''
  }
  return String(value)
}

/**
 * Shallow-clone `data` and coerce keys listed in `schema`.
 * Unknown keys are copied as-is.
 */
export function coercePayload<T extends Record<string, unknown>>(
  data: T,
  schema: CoerceSchema
): Record<string, unknown> {
  const out: Record<string, unknown> = { ...data }
  for (const [key, kind] of Object.entries(schema)) {
    if (!(key in out) && data[key as keyof T] === undefined) {
      continue
    }
    const raw = out[key]
    switch (kind) {
      case 'number':
        out[key] = toNumber(raw)
        break
      case 'boolean':
        out[key] = toBoolean(raw)
        break
      case 'date': {
        const d = toDate(raw)
        if (d === null) {
          delete out[key]
        } else {
          out[key] = d
        }
        break
      }
      case 'string':
        out[key] = toString(raw)
        break
    }
  }
  return out
}

/** Employee profile PUT body (matches Go EmployeeProfile JSON tags). */
export const employeeProfileSchema: CoerceSchema = {
  pregnancyWeeks: 'number',
  age: 'number',
  height: 'number',
  weight: 'number',
  isPregnant: 'boolean',
  hasMedicalHistory: 'boolean',
  hasHypertension: 'boolean',
  hasDiabetes: 'boolean',
  hasHyperlipidemia: 'boolean',
  hasHyperuricemia: 'boolean',
  hasCoronaryHeartDisease: 'boolean',
  hasStroke: 'boolean',
  hasOtherCardiovascularDiseases: 'boolean',
  hasAsthma: 'boolean',
  hasCOPD: 'boolean',
  hasPepticUlcer: 'boolean',
  hasMalignantTumor: 'boolean',
  hasLungCancer: 'boolean',
  hasOtherCancer: 'boolean',
  hasSevereMentalDisorders: 'boolean',
  hasTuberculosis: 'boolean',
  hasHepatitis: 'boolean',
  hasOccupationalDisease: 'boolean',
  hasChronicKidneyDisease: 'boolean',
  hasChronicLiverDisease: 'boolean',
  hasImmunodeficiency: 'boolean',
  hasTyphus: 'boolean',
  isPostpartumInSixWeeks: 'boolean',
  hasDustExposure: 'boolean',
  hasOtherDiseases: 'boolean',
  isVaccinatedForCOVID: 'boolean',
  isVaccinatedForFlu: 'boolean',
  isVaccinatedForPlague: 'boolean',
  isVaccinatedForBCG: 'boolean',
  isVaccinatedForHepatitis: 'boolean',
  birthDate: 'date',
  workOnPlateauStartDate: 'date',
  gender: 'string',
  department: 'string',
  ethnicity: 'string',
  educationLevel: 'string',
  idNumber: 'string',
  name: 'string',
}

/** Work environment POST body. */
export const workEnvironmentSchema: CoerceSchema = {
  avgActivityFaceToFace: 'number',
  avgActivityPhysicalContact: 'number',
  avgActivityRoomPeople: 'number',
  avgBreakfastFaceToFace: 'number',
  avgBreakfastPhysicalContact: 'number',
  avgBreakfastRoomPeople: 'number',
  avgDinnerFaceToFace: 'number',
  avgDinnerPhysicalContact: 'number',
  avgDinnerRoomPeople: 'number',
  avgLunchFaceToFace: 'number',
  avgLunchPhysicalContact: 'number',
  avgLunchRoomPeople: 'number',
  avgMonthlyPressure: 'number',
  avgWorkFaceToFace: 'number',
  avgWorkPhysicalContact: 'number',
  avgWorkRoomPeople: 'number',
  collectiveActivitiesCount: 'number',
  dormitoryArea: 'number',
  dormitoryBuildingNumber: 'number',
  dormitoryOccupants: 'number',
  dormitoryRoomNumber: 'number',
  maxTemperature: 'number',
  minTemperature: 'number',
  windowOpenArea: 'number',
  workStationAltitude: 'number',
  ventilationCondition: 'string',
  workStationName: 'string',
  submissionTime: 'date',
}

/** Nested numeric fields under diagnosispersonalinfo. */
export const diagnosisPersonalInfoSchema: CoerceSchema = {
  height: 'number',
  weight: 'number',
  bloodSugarFasting: 'number',
  bloodSugarRandom: 'number',
  cholesterolTotal: 'number',
  triglycerides: 'number',
  cholesterolLipid: 'number',
  hdl: 'number',
  ldl: 'number',
  bloodOxygen: 'number',
}

/**
 * Normalize survey form quirks before profile PUT:
 * - pregnant ("true"/"false") → isPregnant (boolean)
 * - hasMedicalHistory may be string
 */
export function normalizeEmployeeProfilePayload(
  form: Record<string, unknown>
): Record<string, unknown> {
  const raw = { ...form }
  if (raw.pregnant !== undefined && raw.isPregnant === undefined) {
    raw.isPregnant = raw.pregnant
  }
  delete raw.pregnant
  // phoneNumber stays on User and is not writable via this form
  delete raw.phoneNumber
  // name is accepted by the Go API and synced onto users.name
  return coercePayload(raw, employeeProfileSchema)
}
