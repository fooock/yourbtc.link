export class Util {
  public static transformBytes(bytes: number): string {
    if (bytes === 0)
      return '0 Bytes';

    const k = 1024;
    const sizes = ['bytes', 'Kb', 'Mb', 'Gb', 'Tb'];

    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
  }
}
